package api

import (
	"CRAZY/middleware"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Login 登录
func Login(c *gin.Context) {
	generateToken(c)
	// var loginReq model.LoginReq
	// if c.BindJSON(&loginReq) == nil {
	// 	isPass, user, err := model.LoginCheck(loginReq)
	// 	if isPass {
	// 		generateToken(c, user)
	// 	} else {
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"status": -1,
	// 			"msg":    "验证失败," + err.Error(),
	// 		})
	// 	}
	// } else {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status": -1,
	// 		"msg":    "json 解析失败",
	// 	})
	// }
}

// 生成令牌
func generateToken(c *gin.Context) {
	j := &middleware.JWT{
		SigningKey: []byte("crazy"),
	}
	claims := middleware.CustomClaims{
		ID:    "xxx",
		Name:  "xxx",
		Phone: "xxxx",
	}

	claims.IssuedAt = time.Now().Unix()                // 签名生效时间
	claims.ExpiresAt = int64(time.Now().Unix() + 3600) // 过期时间 一小时

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	log.Println(token)
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"data":   token,
	})
	return
}

func GetNumber(c *gin.Context) {
	p := c.Query("p")
	c.JSON(http.StatusOK, gin.H{
		"status": p,
		"msg":    "登录成功！",
	})
	return
}