package api

import (
	"CRAZY/utils"
	"bytes"
	"net/http"
	"time"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

func GetCaptcha(c *gin.Context) {
	captchaID := captcha.NewLen(4)

	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")
	c.Writer.Header().Set("Content-Type", "image/png")
	c.Writer.Header().Set("Captcha-ID", captchaID)

	var content bytes.Buffer
	captcha.WriteImage(&content, captchaID, 100, 40)
	http.ServeContent(c.Writer, c.Request, captchaID+".png", time.Time{}, bytes.NewReader(content.Bytes()))
}

type CaptchaForm struct {
	CaptchaID string `form:"captchaID" binding:"required"`
	Code      string `form:"code" binding:"required"`
}

func PostCaptcha(c *gin.Context) {
	var form CaptchaForm
	err := c.ShouldBind(&form)

	if err == nil {
		if captcha.VerifyString(form.CaptchaID, form.Code) {
			utils.OkDetailed("验证成功", "success", c)
		} else {
			utils.FailWithMessage("验证码错误", c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}
