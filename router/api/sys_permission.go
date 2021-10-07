package api

import (
	"CRAZY/model"
	sysPermissionService "CRAZY/services/sys_permission"
	"CRAZY/utils"

	"github.com/gin-gonic/gin"
)

type PermissionForm struct {
	Name   string `form:"name" binding:"required"`
	Key    string `form:"key" binding:"required"`
	Status int    `form:"status" binding:"required"`
	PID    int    `form:"pid" binding:"required"`
}

// PostPermission 新增权限
func PostPermission(c *gin.Context) {
	var form PermissionForm
	err := c.ShouldBind(&form)
	if err == nil {
		Model := &model.Permission{
			Name:   form.Name,
			Key:    form.Key,
			Status: form.Status,
			PID:    uint(form.PID),
		}
		res, resErr := sysPermissionService.Create(Model)
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
	id := utils.StrToUInt(c.Param("id"))
	resErr := sysPermissionService.DeleteById(id)
	if resErr == nil {
		utils.Ok(c)
	} else {
		utils.FailWithMessage(resErr.Error(), c)
	}
}

// PutPermissionById 修改权限
func PutPermissionById(c *gin.Context) {
	id := utils.StrToUInt(c.Param("id"))
	var form PermissionForm
	err := c.ShouldBind(&form)
	if err == nil {
		Model := &model.Permission{
			Name:   form.Name,
			Key:    form.Key,
			Status: form.Status,
		}
		res, resErr := sysPermissionService.UpdateById(id, Model)
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
	id := utils.StrToUInt(c.Param("id"))
	res := sysPermissionService.GetById(id)
	utils.OkDetailed(res, "success", c)
}

// GetPermission 获取权限列表树
func GetPermissionTreeById(c *gin.Context) {
	id := utils.StrToUInt(c.Param("id"))
	res := sysPermissionService.GetTree(uint(id))
	utils.OkDetailed(res, "success", c)
}
