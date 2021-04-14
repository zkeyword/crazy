package api

import (
	"CRAZY/services"
	"CRAZY/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

// PostRole 新增角色
func PostRole(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := services.NewRoleService.Get(id)
	utils.OkDetailed(gin.H{
		"user": user,
	}, "success", c)
}

// DelRoleById 删除角色
func DelRoleById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := services.NewRoleService.Get(id)
	utils.OkDetailed(gin.H{
		"user": user,
	}, "success", c)
}

// PutRoleById 修改角色
func PutRoleById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := services.NewRoleService.Get(id)
	utils.OkDetailed(gin.H{
		"user": user,
	}, "success", c)
}

// GetRoleById 获取角色
func GetRoleById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	user := services.NewRoleService.Get(id)
	utils.OkDetailed(gin.H{
		"user": user,
	}, "success", c)
}
