package captcha

import (
	"fmt"
	"github.com/mojocn/base64Captcha"
)

type configJsonBody struct {
	Id            string
	CaptchaType   string
	VerifyValue   string
	DriverAudio   *base64Captcha.DriverAudio
	DriverString  *base64Captcha.DriverString
	DriverChinese *base64Captcha.DriverChinese
	DriverMath    *base64Captcha.DriverMath
	DriverDigit   *base64Captcha.DriverDigit
}

var store = base64Captcha.DefaultMemStore

//GenerateCaptcha 生成验证码，返回验证码和base64字符串
func GenerateCaptcha() {

	// var driver base64Captcha.Driver
	var param configJsonBody

	// switch param.CaptchaType {
	// case "audio":
	// 	driver = param.DriverAudio
	// case "string":
	// 	driver = param.DriverString.ConvertFonts()
	// case "math":
	// 	driver = param.DriverMath.ConvertFonts()
	// case "chinese":
	// 	driver = param.DriverChinese.ConvertFonts()
	// default:
	// 	driver = param.DriverDigit
	// }

	fmt.Println("11")

	c := base64Captcha.NewCaptcha(param.DriverDigit, store)

	fmt.Println(2222, c)
	if id, b64s, err := c.Generate(); err != nil {
		fmt.Println(id, b64s, err)
	}

	return
}

// VerifyCaptcha 校验
func VerifyCaptcha(idKey, verifyValue string) bool {
	verifyResult := store.Verify(idKey, verifyValue, true)
	if verifyResult {
		return true
	}
	return false
}
