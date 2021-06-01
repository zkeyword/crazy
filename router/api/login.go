package api

import (
	"CRAZY/middleware"
	"CRAZY/services"
	"CRAZY/utils"
	"CRAZY/utils/xor"
	"time"

	"github.com/gin-gonic/gin"
)

type LoginUserForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// Login 登录
func Login(c *gin.Context) {
	var form LoginUserForm
	err := c.ShouldBind(&form)
	if err == nil {
		res, resErr := services.NewUserService.GetByUserName(form.Username)
		if resErr == nil {
			if xor.Enc(form.Password) == res.Password {
				j := &middleware.JWT{
					SigningKey: []byte("crazy"),
				}
				claims := middleware.CustomClaims{
					Name: form.Username,
				}

				claims.IssuedAt = time.Now().Unix()                // 签名生效时间
				claims.ExpiresAt = int64(time.Now().Unix() + 3600) // 过期时间 一小时
				token, jwtErr := j.CreateToken(claims)
				if jwtErr == nil {
					utils.OkDetailed(token, "success", c)
				}
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
