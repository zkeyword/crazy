package core

import (
	// "CRAZY/controllers"

	"github.com/gin-gonic/gin"
)

// Routers 总路由
func Routers() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}
