package api

import (
	"CRAZY/utils"

	"github.com/gin-gonic/gin"
)

// GetTags 测试
func GetTags(c *gin.Context) {
	utils.OkDetailed(gin.H{
		"message": "pong",
	}, "请求成功", c)
	// c.JSON(200, gin.H{
	// 	"message": "pong",
	// })
}
