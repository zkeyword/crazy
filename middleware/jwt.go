package middleware

import (
	"CRAZY/config"
	"CRAZY/utils"
	"CRAZY/utils/xor"
	"errors"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWTAuth 中间件，检查token
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			utils.FailWithMessage("请求未携带token，无权限访问", c)
			c.Abort()
			return
		}
		j := NewJWT()
		// parseToken 解析token包含的信息
		tokenString := strings.TrimPrefix(token, "Bearer ")
		claims, err := j.ParseToken(tokenString)
		if err != nil {
			if err == ErrTokenExpired {
				utils.FailWithMessage("授权已过期", c)
				c.Abort()
				return
			}
			utils.FailWithMessage(err.Error(), c)
			c.Abort()
			return
		}
		// if claims.UserID == 0 {
		// 	utils.FailWithMessage("Authorization出错, 清重新登录", c)
		// 	c.Abort()
		// 	return
		// }
		// 继续交由下一个路由处理,并将解析出的信息传递下去
		c.Set("userID", claims.UserID)
		c.Set("permissionKeys", xor.Dec(claims.PermissionKeys))
	}
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte
}

// 一些常量
var (
	ErrTokenExpired     error  = errors.New("token is expired")
	ErrTokenNotValidYet error  = errors.New("token not active yet")
	ErrTokenMalformed   error  = errors.New("that's not even a token")
	ErrTokenInvalid     error  = errors.New("couldn't handle this token")
	SignKey             string = "crazy"
)

// 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	UserName       string `json:"userName"`
	UserID         uint   `json:"userID"`
	PermissionKeys string `json:"permissionKeys"`
	jwt.StandardClaims
}

// 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		[]byte(GetSignKey()),
	}
}

// 获取signKey
func GetSignKey() string {
	return SignKey
}

// SetSignKey 这是SignKey
func SetSignKey(key string) string {
	SignKey = key
	return SignKey
}

// CreateToken 生成一个token
func (j *JWT) CreateToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// ParseToken 解析Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, ErrTokenInvalid
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.StandardClaims.ExpiresAt = config.JWTExpiresAt
		return j.CreateToken(*claims)
	}
	return "", ErrTokenInvalid
}
