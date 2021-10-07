package api

import (
	"CRAZY/model"
	sysRoleService "CRAZY/services/sys_role"
	"CRAZY/utils"
	"html"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleForm struct {
	Name           string `form:"name" binding:"required"`
	PermissionKeys string `form:"permission" binding:"required"`
}

// PostRole 新增角色
func PostRole(c *gin.Context) {
	var form RoleForm
	err := c.ShouldBind(&form)
	if err == nil {
		Model := &model.Role{
			Name: form.Name,
		}
		res, resErr := sysRoleService.Create(Model, form.PermissionKeys)
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
	id := utils.StrToUInt(c.Param("id"))
	resErr := sysRoleService.DeleteById(id)
	if resErr == nil {
		utils.Ok(c)
	} else {
		utils.FailWithMessage(resErr.Error(), c)
	}
}

// PutRoleById 修改角色
func PutRoleById(c *gin.Context) {
	id := utils.StrToUInt(c.Param("id"))
	var form RoleForm
	err := c.ShouldBind(&form)
	if err == nil {
		Model := &model.Role{
			Name: form.Name,
		}
		res, resErr := sysRoleService.UpdateById(id, Model, form.PermissionKeys)
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
	id := utils.StrToUInt(c.Param("id"))
	res := sysRoleService.GetById(id)
	utils.OkDetailed(res, "success", c)
}

// GetRole 获取角色
func GetRole(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	name := html.EscapeString(c.Query("name"))
	res, _ := sysRoleService.Get(page, pageSize, name)
	utils.OkDetailed(res, "success", c)
}

// GetRolePermissionByRoleID 获取角色权限
func GetRolePermissionByRoleID(c *gin.Context) {
	id := utils.StrToUInt(c.Param("id"))
	res := sysRoleService.GetRolePermissionByRoleID(id)
	utils.OkDetailed(res, "success", c)
}

type RolePermissionForm struct {
	PermissionKeys string `form:"permission" binding:"required"`
}

// PostRolePermissionByRoleID 修改角色权限
func PostRolePermissionByRoleID(c *gin.Context) {
	var form RolePermissionForm
	err := c.ShouldBind(&form)
	if err == nil {
		id := utils.StrToUInt(c.Param("id"))
		res, resErr := sysRoleService.PostRolePermissionByRoleID(id, form.PermissionKeys)
		if resErr == nil {
			utils.OkDetailed(res, "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

// GetRoleUserByRoleID 获取角色关联的用户
func GetRoleUserByRoleID(c *gin.Context) {
	id := utils.StrToUInt(c.Param("id"))
	res := sysRoleService.GetRoleUserByRoleID(id)
	utils.OkDetailed(res, "success", c)
}

type PostRoleUserByRoleIDForm struct {
	UserID   int    `form:"userId" binding:"required"`
	Username string `form:"username" binding:"required"`
}

// PostRoleUserByRoleID 新增角色关联用户
func PostRoleUserByRoleID(c *gin.Context) {
	id := utils.StrToUInt(c.Param("id"))
	var form PostRoleUserByRoleIDForm
	err := c.ShouldBind(&form)
	if err == nil {
		res, resErr := sysRoleService.PostRoleUserByRoleID(uint(form.UserID), form.Username, id)
		if resErr == nil {
			utils.OkDetailed(res, "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}

type DelectRoleUserByRoleIAndUserIDForm struct {
	UserID int `form:"userId" binding:"required"`
}

// DelectRoleUserByRoleIAndUserID 删除角色关联的用户
func DelectRoleUserByRoleIAndUserID(c *gin.Context) {
	id := utils.StrToUInt(c.Param("id"))
	var form PostRoleUserByRoleIDForm
	err := c.ShouldBind(&form)
	if err == nil {
		resErr := sysRoleService.DeleteByRoleIdAndUserId(uint(form.UserID), id)
		if resErr == nil {
			utils.OkDetailed("删除成功！", "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	} else {
		utils.FailWithMessage(err.Error(), c)
	}
}
