package middleware

import (
	"CRAZY/config"
	"CRAZY/utils"
	"CRAZY/utils/db"
	"CRAZY/utils/xor"
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

/* ---------- 常量 ---------- */
const (
	authHeader         = "Authorization"
	bearerPrefix       = "Bearer "
	blacklistPrefix    = "auth:blacklist:jti:"
	globalLogoutPrefix = "auth:user:global_logout:"
)

/* ---------- 错误定义 ---------- */
var (
	ErrMissingHeader    = errors.New("missing Authorization header")
	ErrInvalidBearer    = errors.New("token must start with 'Bearer '")
	ErrTokenExpired     = errors.New("token is expired")
	ErrTokenNotValidYet = errors.New("token not active yet")
	ErrTokenMalformed   = errors.New("token malformed")
	ErrTokenInvalid     = errors.New("token invalid")
	ErrSessionKicked    = errors.New("session expired by global logout")
	ErrTokenRevoked     = errors.New("token has been revoked")
)

/* ---------- JWT 结构 ---------- */
type JWT struct{ signingKey []byte }

func NewJWT() *JWT { return &JWT{signingKey: []byte(config.JWTSignKey)} }

/* ---------- 自定义 Claims ---------- */
type CustomClaims struct {
	UserName       string `json:"userName"`
	UserID         uint   `json:"userID"`
	PermissionKeys string `json:"permissionKeys"`
	jwt.StandardClaims
}

/* ---------- 中间件入口 ---------- */
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, raw, err := extractAndValidate(c)
		if err != nil {
			logAuthErr(c, err) // 统一日志
			utils.FailWithUnauthorized(friendlyMsg(err), c)
			c.Abort()
			return
		}

		// 注入上下文
		c.Set("jwtToken", raw)
		c.Set("userID", claims.UserID)
		c.Set("userName", claims.UserName)
		c.Set("permissionKeys", xor.Dec(claims.PermissionKeys))
		c.Next()
	}
}

/* ---------- 主流程 ---------- */
func extractAndValidate(c *gin.Context) (*CustomClaims, string, error) {
	raw, err := getRawToken(c)
	if err != nil {
		return nil, "", err
	}

	claims, err := NewJWT().ParseToken(raw)
	if err != nil {
		return nil, raw, mapParseError(err)
	}

	if checkGlobalLogout(claims.UserID, claims.IssuedAt) {
		return nil, raw, ErrSessionKicked
	}
	if checkBlacklist(claims.Id) {
		return nil, raw, ErrTokenRevoked
	}
	return claims, raw, nil
}

/* ---------- 工具函数 ---------- */
func getRawToken(c *gin.Context) (string, error) {
	auth := c.GetHeader(authHeader)
	if auth == "" {
		return "", ErrMissingHeader
	}
	if !strings.HasPrefix(auth, bearerPrefix) {
		return "", ErrInvalidBearer
	}
	return strings.TrimPrefix(auth, bearerPrefix), nil
}

func mapParseError(err error) error {
	if ve, ok := err.(*jwt.ValidationError); ok {
		switch {
		case ve.Errors&jwt.ValidationErrorMalformed != 0:
			return ErrTokenMalformed
		case ve.Errors&jwt.ValidationErrorExpired != 0:
			return ErrTokenExpired
		case ve.Errors&jwt.ValidationErrorNotValidYet != 0:
			return ErrTokenNotValidYet
		}
	}
	return ErrTokenInvalid
}

func checkGlobalLogout(userID uint, issuedAt int64) bool {
	val, _ := db.GetKey(globalLogoutPrefix + strconv.FormatUint(uint64(userID), 10))
	if val == "" {
		return false
	}
	logoutTime, _ := strconv.ParseInt(val, 10, 64)
	return issuedAt <= logoutTime
}

func checkBlacklist(jti string) bool {
	ok, _ := db.Exists(blacklistPrefix + jti)
	return ok
}

/* ---------- 日志 & 友好提示 ---------- */
func logAuthErr(c *gin.Context, err error) {
	utils.Log.Errorf("[JWTAuth] ip=%s path=%s error=%v", c.ClientIP(), c.Request.URL.Path, err)
}

func friendlyMsg(err error) string {
	switch {
	case errors.Is(err, ErrMissingHeader):
		return "请先登录"
	case errors.Is(err, ErrInvalidBearer):
		return "Authorization 格式错误（需 Bearer <token>）"
	case errors.Is(err, ErrTokenExpired):
		return "登录已过期，请重新登录"
	case errors.Is(err, ErrTokenMalformed):
		return "Token 格式非法"
	case errors.Is(err, ErrTokenNotValidYet):
		return "Token 尚未生效"
	case errors.Is(err, ErrSessionKicked):
		return "账号已在其他设备登出，请重新登录"
	case errors.Is(err, ErrTokenRevoked):
		return "当前会话已被终止"
	default:
		return "Token 验证失败"
	}
}

/* ---------- JWT 方法 ---------- */
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.signingKey)
}

func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.signingKey, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

// RefreshToken 跳过过期校验，仅验证签名 + 黑名单
func (j *JWT) RefreshToken(oldToken string) (string, error) {
	claims, err := j.ParseTokenForRefresh(oldToken)
	if err != nil {
		return "", err
	}
	// 过期超过 7 天不允许刷新
	if time.Now().Unix()-claims.ExpiresAt > 7*24*3600 {
		return "", errors.New("token too old, please login again")
	}

	newClaims := *claims
	newClaims.Id = uuid.NewString()
	newClaims.IssuedAt = time.Now().Unix()
	newClaims.ExpiresAt = config.JWTExpiresAt

	return j.CreateToken(newClaims)
}

func (j *JWT) ParseTokenForRefresh(tokenString string) (*CustomClaims, error) {
	parser := &jwt.Parser{SkipClaimsValidation: true}
	token, err := parser.ParseWithClaims(tokenString, &CustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return j.signingKey, nil
	})
	if err != nil {
		return nil, ErrTokenInvalid
	}
	if claims, ok := token.Claims.(*CustomClaims); ok {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}
