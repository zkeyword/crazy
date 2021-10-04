package api

import (
	"CRAZY/utils"
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func GetCaptcha(c *gin.Context) {
	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")
	c.Writer.Header().Set("Content-Type", "image/png")

	captchaID := captcha.NewLen(4)

	session := sessions.Default(c)
	session.Set("captcha", captchaID)
	session.Save()

	var content bytes.Buffer
	captcha.WriteImage(&content, captchaID, 100, 50)
	http.ServeContent(c.Writer, c.Request, captchaID+".png", time.Time{}, bytes.NewReader(content.Bytes()))
}

type CaptchaForm struct {
	Code string `form:"code" binding:"required"`
}

func PostCaptcha(c *gin.Context) {
	var form CaptchaForm
	err := c.ShouldBind(&form)
	session := sessions.Default(c)

	if err == nil {

		fmt.Println(session.Get("captcha"), form.Code)

		if session.Get("captcha") != form.Code {
			utils.FailWithMessage("验证码错误", c)
		} else {
			utils.OkDetailed("验证成功", "success", c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}
