package common

import (
	"CRAZY/utils"
	"CRAZY/utils/xor"
	"bytes"
	"net/http"
	"regexp"

	"github.com/dchest/captcha"
	"github.com/gin-gonic/gin"
)

func GetCaptcha(c *gin.Context) {
	captchaID := captcha.NewLen(4)
	v := c.Query("v")

	if !utils.IsTimestamp(v) {
		utils.FailWithMessage("缺少参数", c)
		return
	}

	c.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Writer.Header().Set("Pragma", "no-cache")
	c.Writer.Header().Set("Expires", "0")
	c.Writer.Header().Set("Content-Type", "image/png")
	c.Writer.Header().Set("Captcha-ID", xor.Enc(captchaID+v))

	var content bytes.Buffer
	captcha.WriteImage(&content, captchaID, 100, 40)
	c.Data(http.StatusOK, "image/png", content.Bytes())
}

type CaptchaForm struct {
	CaptchaID string `form:"captchaID" binding:"required"`
	Time      string `form:"time" binding:"required"`
	Code      string `form:"code" binding:"required"`
}

func PostCaptcha(c *gin.Context) {
	var form CaptchaForm
	err := c.ShouldBind(&form)
	if err == nil {
		str := xor.Dec(form.CaptchaID)
		re := regexp.MustCompile(`^(.*)(\d{13})$`)
		matches := re.FindStringSubmatch(str)
		if form.Time == matches[2] && captcha.VerifyString(matches[1], form.Code) {
			utils.OkDetailed("验证成功", "success", c)
		} else {
			utils.FailWithMessage("验证码错误", c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}
