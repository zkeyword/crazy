package api

import (
	"CRAZY/services"
	"CRAZY/utils"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetUser 测试
func GetUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	fmt.Println(id, err)
	if err == nil {
		fmt.Println(id)
		user, r := services.NewUserService.FindByID(id)
		fmt.Println(user)
		fmt.Println(r)
	}
	utils.OkDetailed(gin.H{
		"message": "pong",
	}, "请求成功", c)
	// c.JSON(200, gin.H{
	// 	"message": "pong",
	// })
}
