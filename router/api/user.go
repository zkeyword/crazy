package api

import (
	"CRAZY/services"
	"CRAZY/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PostUser 新增用户
func PostUser(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := services.NewUserService.Get(id)
	utils.OkDetailed(gin.H{
		"user": user,
	}, "success", c)
}

// DelUserById 删除用户
func DelUserById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := services.NewUserService.Get(id)
	utils.OkDetailed(gin.H{
		"user": user,
	}, "success", c)
}

// PutUserById 修改用户
func PutUserById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := services.NewUserService.Get(id)
	utils.OkDetailed(gin.H{
		"user": user,
	}, "success", c)
}

// GetUserById 获取用户
func GetUserById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := services.NewUserService.Get(id)
	utils.OkDetailed(gin.H{
		"user": user,
	}, "success", c)
}
