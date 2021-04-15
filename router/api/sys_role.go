package api

import (
	"CRAZY/model"
	"CRAZY/services"
	"CRAZY/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleForm struct {
	Name string `form:"name" binding:"required"`
}

// PostRole 新增角色
func PostRole(c *gin.Context) {
	var form RoleForm
	err := c.ShouldBind(&form)
	if err == nil {
		Model := &model.Role{
			Name: form.Name,
		}
		res, resErr := services.NewRoleService.Create(Model)
		if resErr == nil {
			utils.OkDetailed(res, "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

// DelRoleById 删除角色
func DelRoleById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	resErr := services.NewRoleService.DeleteById(id)
	if resErr == nil {
		utils.Ok(c)
	} else {
		utils.FailWithMessage(resErr.Error(), c)
	}
}

// PutRoleById 修改角色
func PutRoleById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var form RoleForm
	err := c.ShouldBind(&form)
	if err == nil {
		Model := &model.Role{
			Name: form.Name,
		}
		res, resErr := services.NewRoleService.UpdateById(id, Model)
		if resErr == nil {
			utils.OkDetailed(res, "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

// GetRoleById 获取角色
func GetRoleById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := services.NewRoleService.Get(id)
	utils.OkDetailed(res, "success", c)
}
