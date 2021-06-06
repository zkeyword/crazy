package api

import (
	"CRAZY/middleware"
	"CRAZY/model"
	"CRAZY/services"
	"CRAZY/utils"
	"CRAZY/utils/xor"
	"html"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginUserForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

type ReturnLoginUser struct {
	ID        uint      `json:"id"`
	Username  string    `json:"username"`
	Status    int       `json:"status"`
	Level     int       `json:"level"`
	ParentID  uint      `json:"parentID"`
	UpdatedAt time.Time `json:"updatedAt"`
	Token     string    `json:"token"`
}

// getToken 获取 JWT token
func getToken(Name string) string {
	j := &middleware.JWT{
		SigningKey: []byte("crazy"),
	}
	claims := middleware.CustomClaims{
		Name: Name,
	}

	claims.IssuedAt = time.Now().Unix()                // 签名生效时间
	claims.ExpiresAt = int64(time.Now().Unix() + 3600) // 过期时间 一小时
	token, _ := j.CreateToken(claims)
	return token
}

// Login 登录
func Login(c *gin.Context) {
	var form LoginUserForm
	err := c.ShouldBind(&form)
	if err == nil {
		res, resErr := services.NewUserService.GetByUserName(html.EscapeString(form.Username))
		if resErr == nil {
			if xor.Enc(form.Password) == res.Password {
				user := &ReturnLoginUser{
					ID:        res.ID,
					Username:  res.Username,
					Status:    res.Status,
					Level:     res.Level,
					ParentID:  res.ParentID,
					UpdatedAt: res.UpdatedAt,
					Token:     getToken(form.Username),
				}
				utils.OkDetailed(user, "success", c)
			} else {
				utils.FailWithMessage("密码错误", c)
			}
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

// Register 注册
func Register(c *gin.Context) {
	var form LoginUserForm
	err := c.ShouldBind(&form)
	if err == nil {
		Model := &model.User{
			Username: html.EscapeString(form.Username),
			Password: xor.Enc(form.Password),
			Status:   1,
		}
		res, resErr := services.NewUserService.Create(Model, "999")
		if resErr == nil {
			user := &ReturnLoginUser{
				ID:        res.ID,
				Username:  res.Username,
				Status:    res.Status,
				Level:     res.Level,
				ParentID:  res.ParentID,
				UpdatedAt: res.UpdatedAt,
				Token:     getToken(form.Username),
			}
			utils.OkDetailed(user, "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}
