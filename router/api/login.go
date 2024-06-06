package api

import (
	"CRAZY/config"
	"CRAZY/middleware"
	"CRAZY/model"
	sysUserService "CRAZY/services/sys_user"
	"CRAZY/utils"
	"CRAZY/utils/xor"
	"html"
	"regexp"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
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
	Level          int       `json:"level"`
	ParentID       uint      `json:"parentID"`
	UpdatedAt      time.Time `json:"updatedAt"`
	Token          string    `json:"token"`
	PermissionKeys string    `json:"permissions"`
	UserKey        string    `json:"userKey"`
}

// getToken 获取 JWT token
func getToken(UserName string, UserID uint, PermissionKeys string) string {
	j := &middleware.JWT{
		SigningKey: []byte("crazy"),
	}
	claims := middleware.CustomClaims{
		UserName:       UserName,
		UserID:         UserID,
		PermissionKeys: xor.Enc(PermissionKeys),
	}

	claims.IssuedAt = time.Now().Unix()    // 签名生效时间
	claims.ExpiresAt = config.JWTExpiresAt // 过期时间 30D
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

		// 用户是否存在
		res, resErr := sysUserService.GetByUserName(html.EscapeString(form.Username))
		if resErr != nil {
			utils.FailWithMessage(resErr.Error(), c)
			return
		}

		// 校验密码
		if xor.Enc(form.Password) != res.Password {
			utils.FailWithMessage("密码错误", c)
			return
		}

		userDetailRes := sysUserService.GetUserRolePermissionByUserId(res.ID)
		userKey := utils.StringWithCharset(5)
		user := &ReturnLoginUser{
			ID:             res.ID,
			Username:       res.Username,
			Status:         res.Status,
			Level:          res.Level,
			ParentID:       res.ParentID,
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
		Model := &model.User{
			Username: html.EscapeString(form.Username),
			Password: xor.Enc(form.Password),
			Status:   1,
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
			Level:     res.Level,
			ParentID:  res.ParentID,
			UpdatedAt: res.UpdatedAt,
			Token:     getToken(form.Username, res.ID, ""),
		}
		utils.OkDetailed(user, "success", c)
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}
