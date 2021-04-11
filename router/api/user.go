package api

import (
	"CRAZY/utils"

	"github.com/gin-gonic/gin"
)

// GetUser 测试
func GetUser(c *gin.Context) {
	utils.OkDetailed(gin.H{
		"message": "pong",
	}, "请求成功", c)
	// c.JSON(200, gin.H{
	// 	"message": "pong",
	// })
}
