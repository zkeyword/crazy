package middleware

import (
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
)

// Cors 处理跨域请求,支持options访问
func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {

		origin := c.Request.Header.Get("Origin")

		reg := regexp.MustCompile(`localhost*|127.0.0.1*|0.0.0.0*`)
		if len(reg.FindAllString(origin, -1)) > 0 {
			c.Header("Access-Control-Allow-Origin", origin)
		}

		c.Header("Access-Control-Allow-Headers", "Content-Type, AccessToken, X-CSRF-Token, x-requested-with, Authorization, Token, authKey")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type")
		c.Header("Access-Control-Allow-Credentials", "true")

		// 放行所有OPTIONS方法
		method := c.Request.Method
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}

		// 处理请求
		c.Next()
	}
}
