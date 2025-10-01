package sys

import (
	"CRAZY/config"
	"CRAZY/middleware"
	"CRAZY/model/sys"
	sysUserService "CRAZY/services/sys/sys_user"
	"CRAZY/utils"
	"CRAZY/utils/db"
	"CRAZY/utils/xor"
	"fmt"
	"html"
	"regexp"
	"strconv"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LoginUserForm struct {
	Username  string `form:"username" binding:"required"`
	Password  string `form:"password" binding:"required"`
	CaptchaID string `form:"captchaID" binding:"required"`
	Code      string `form:"code" binding:"required"`
	Time      string `form:"time" binding:"required"`
}

type ReturnLoginUser struct {
	ID             uint      `json:"id"`
	Username       string    `json:"username"`
	Status         int       `json:"status"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Token          string    `json:"token"`
	PermissionKeys string    `json:"permissions"`
	UserKey        string    `json:"userKey"`
}

// getToken 获取 JWT token
func getToken(UserName string, UserID uint, PermissionKeys string) string {
	j := middleware.NewJWT()
	claims := middleware.CustomClaims{
		UserName:       UserName,
		UserID:         UserID,
		PermissionKeys: xor.Enc(PermissionKeys),
	}

	claims.IssuedAt = time.Now().Unix()    // 签名生效时间
	claims.ExpiresAt = config.JWTExpiresAt // 过期时间 30D
	claims.Id = uuid.NewString()           // 唯一标识
	token, _ := j.CreateToken(claims)
	return token
}

// Login 登录
func Login(c *gin.Context) {
	var form LoginUserForm
	err := c.ShouldBind(&form)
	if err == nil {
		// 校验验证码
		str := xor.Dec(form.CaptchaID)
		re := regexp.MustCompile(`^(.*)(\d{13})$`)
		matches := re.FindStringSubmatch(str)
		if !(form.Time == matches[2] && captcha.VerifyString(matches[1], form.Code)) {
			utils.FailWithMessage("验证码错误", c)
			return
		}

		res, resErr := sysUserService.GetByUserName(html.EscapeString(form.Username))
		if resErr != nil || xor.Enc(form.Password) != res.Password {
			utils.FailWithMessage("用户名或者密码错误", c)
			return
		}

		if res.Status == -1 {
			utils.FailWithMessage("该账户已被禁用", c)
			return
		}

		// Model := &sys.SysUser{
		// 	LoginStatus: 1,
		// }
		// sysUserService.PutUserById(res.ID, Model, "")
		userDetailRes := sysUserService.GetUserRolePermissionByUserId(res.ID)
		userKey := utils.StringWithCharset(5)
		user := &ReturnLoginUser{
			ID:             res.ID,
			Username:       res.Username,
			Status:         res.Status,
			UpdatedAt:      res.UpdatedAt,
			Token:          getToken(form.Username, res.ID, userDetailRes.PermissionKeys),
			PermissionKeys: xor.XorEncryptDecrypt(userDetailRes.PermissionKeys, userKey),
			UserKey:        userKey,
		}
		utils.OkDetailed(user, "success", c)
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

// Register 注册
func Register(c *gin.Context) {
	var form LoginUserForm
	err := c.ShouldBind(&form)
	if err == nil {
		// 校验验证码
		str := xor.Dec(form.CaptchaID)
		re := regexp.MustCompile(`^(.*)(\d{13})$`)
		matches := re.FindStringSubmatch(str)
		if !(form.Time == matches[2] && captcha.VerifyString(matches[1], form.Code)) {
			utils.FailWithMessage("验证码错误", c)
			return
		}

		// 创建用户并默认999角色
		Model := &sys.SysUser{
			Username:    html.EscapeString(form.Username),
			Password:    xor.Enc(form.Password),
			Status:      1,
			LoginStatus: -1,
		}
		res, resErr := sysUserService.Create(Model, "999")
		if resErr != nil {
			utils.FailWithMessage(resErr.Error(), c)
			return
		}

		user := &ReturnLoginUser{
			ID:        res.ID,
			Username:  res.Username,
			Status:    res.Status,
			UpdatedAt: res.UpdatedAt,
			Token:     getToken(form.Username, res.ID, ""),
		}
		utils.OkDetailed(user, "success", c)
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

func Logout(c *gin.Context) {
	tokenVal, exists := c.Get("jwtToken")
	if !exists {
		utils.FailWithMessage("未获取到有效登录凭证", c)
		return
	}

	token, ok := tokenVal.(string)
	if !ok {
		utils.FailWithMessage("token 格式异常", c)
		return
	}

	j := middleware.NewJWT()
	claims, err := j.ParseTokenForRefresh(token)
	if err != nil {
		utils.OkDetailed("退出成功", "success", c)
		return
	}

	remaining := claims.ExpiresAt - time.Now().Unix()
	ttl := int(remaining)
	if ttl <= 0 {
		ttl = 60
	}
	blacklistKey := "auth:blacklist:jti:" + claims.Id
	_ = db.SetKeyEx(blacklistKey, "1", ttl)

	utils.OkDetailed("退出成功！", "success", c)
}

func LogoutAll(c *gin.Context) {
	userID := c.GetUint("userID")

	now := time.Now().Unix()
	globalLogoutKey := fmt.Sprintf("auth:user:global_logout:%d", userID)
	_ = db.SetKeyEx(globalLogoutKey, strconv.FormatInt(now, 10), 7*24*3600)

	utils.OkDetailed("所有设备已退出", "success", c)
}

type ReturnRefreshToken struct {
	Token string `json:"token"`
}

// Refresh 刷新 token
func Refresh(c *gin.Context) {
	j := middleware.NewJWT()
	jwtToken, _ := c.Get("jwtToken")
	jwtTokenStr, _ := jwtToken.(string)
	refreshToken, _ := j.RefreshToken(jwtTokenStr)
	res := &ReturnRefreshToken{
		Token: refreshToken,
	}
	utils.OkDetailed(res, "success", c)
}
