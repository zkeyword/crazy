package api

import (
	"CRAZY/model"
	"CRAZY/services"
	"CRAZY/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PermissionForm struct {
	Name string `form:"name" binding:"required"`
}

// PostPermission 新增权限
func PostPermission(c *gin.Context) {
	var form PermissionForm
	err := c.ShouldBind(&form)
	if err == nil {
		Model := &model.Permission{
			Name: form.Name,
		}
		res, resErr := services.NewPermissionService.Create(Model)
		if resErr == nil {
			utils.OkDetailed(res, "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

// DelPermissionById 删除权限
func DelPermissionById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	resErr := services.NewPermissionService.DeleteById(id)
	if resErr == nil {
		utils.Ok(c)
	} else {
		utils.FailWithMessage(resErr.Error(), c)
	}
}

// PutPermissionById 修改权限
func PutPermissionById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var form PermissionForm
	err := c.ShouldBind(&form)
	if err == nil {
		Model := &model.Permission{
			Name: form.Name,
		}
		res, resErr := services.NewPermissionService.UpdateById(id, Model)
		if resErr == nil {
			utils.OkDetailed(res, "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

// GetPermissionById 获取权限
func GetPermissionById(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	res := services.NewPermissionService.Get(id)
	utils.OkDetailed(res, "success", c)
}
