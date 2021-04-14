package api

import (
	"CRAZY/services"
	"CRAZY/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PostPermission 新增权限
func PostPermission(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := services.NewPermissionService.Get(id)
	utils.OkDetailed(gin.H{
		"user": user,
	}, "success", c)
}

// DelPermissionById 删除权限
func DelPermissionById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := services.NewPermissionService.Get(id)
	utils.OkDetailed(gin.H{
		"user": user,
	}, "success", c)
}

// PutPermissionById 修改权限
func PutPermissionById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := services.NewPermissionService.Get(id)
	utils.OkDetailed(gin.H{
		"user": user,
	}, "success", c)
}

// GetPermissionById 获取权限
func GetPermissionById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := services.NewPermissionService.Get(id)
	utils.OkDetailed(gin.H{
		"user": user,
	}, "success", c)
}
