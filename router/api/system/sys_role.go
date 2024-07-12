package system

import (
	"CRAZY/model"
	sysRoleService "CRAZY/services/sys/sys_role"
	"CRAZY/utils"
	"html"
	"strconv"

	"github.com/gin-gonic/gin"
)

type RoleForm struct {
	Name           string `form:"name" binding:"required"`
	Desc           string `form:"desc"`
	PermissionKeys string `form:"permissionKeys" binding:"required"`
}

// PostRole 新增角色
func PostRole(c *gin.Context) {
	if utils.CheckPermission(c, "PostRole") {
		var form RoleForm
		err := c.ShouldBind(&form)
		if err == nil {
			Model := &model.Role{
				Name: form.Name,
				Desc: form.Desc,
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
}

// DelRoleById 删除角色
func DelRoleById(c *gin.Context) {
	if utils.CheckPermission(c, "DelRoleById") {
		id := utils.StrToUInt(c.Param("id"))
		resErr := sysRoleService.DeleteById(id)
		if resErr == nil {
			utils.Ok(c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	}
}

// PutRoleById 修改角色
func PutRoleById(c *gin.Context) {
	if utils.CheckPermission(c, "PutRoleById") {
		id := utils.StrToUInt(c.Param("id"))
		var form RoleForm
		err := c.ShouldBind(&form)
		if err == nil {
			Model := &model.Role{
				Name: form.Name,
				Desc: form.Desc,
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
}

// GetRoleById 获取角色
func GetRoleById(c *gin.Context) {
	if utils.CheckPermission(c, "GetRoleById") {
		id := utils.StrToUInt(c.Param("id"))
		res := sysRoleService.GetById(id)
		utils.OkDetailed(res, "success", c)
	}
}

// GetRole 获取角色
func GetRole(c *gin.Context) {
	if utils.CheckPermission(c, "GetRole") {
		page, _ := strconv.Atoi(c.Query("page"))
		pageSize, _ := strconv.Atoi(c.Query("pageSize"))
		name := html.EscapeString(c.Query("name"))
		res, _ := sysRoleService.Get(page, pageSize, name)
		utils.OkDetailed(res, "success", c)
	}
}

// GetRolePermissionByRoleID 获取角色权限
func GetRolePermissionByRoleID(c *gin.Context) {
	if utils.CheckPermission(c, "GetRolePermissionByRoleID") {
		id := utils.StrToUInt(c.Param("id"))
		res := sysRoleService.GetRolePermissionByRoleID(id)
		utils.OkDetailed(res, "success", c)
	}
}

type RolePermissionForm struct {
	PermissionKeys string `form:"permissionKeys" binding:"required"`
}

// PostRolePermissionByRoleID 修改角色权限
func PostRolePermissionByRoleID(c *gin.Context) {
	if utils.CheckPermission(c, "PostRolePermissionByRoleID") {
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
}

// GetRoleUserByRoleID 获取角色关联的用户
func GetRoleUserByRoleID(c *gin.Context) {
	if utils.CheckPermission(c, "GetRoleUserByRoleID") {
		id := utils.StrToUInt(c.Param("id"))
		res := sysRoleService.GetRoleUserByRoleID(id)
		utils.OkDetailed(res, "success", c)
	}
}

// GetRoleUserByUserID 获取用户关联的角色
func GetRoleUserByUserID(c *gin.Context) {
	if utils.CheckPermission(c, "GetRoleUserByUserID") {
		id := utils.StrToUInt(c.Param("id"))
		res := sysRoleService.GetRoleUserByUserID(id)
		utils.OkDetailed(res, "success", c)
	}
}

type PostRoleUserForm struct {
	Username string `form:"username" binding:"required"`
}

// PostRoleUser 新增角色关联用户
func PostRoleUserByRoleIDAndUserID(c *gin.Context) {
	if utils.CheckPermission(c, "PostRoleUserByRoleIDAndUserID") {
		id := utils.StrToUInt(c.Param("id"))
		userId := utils.StrToUInt(c.Param("userId"))
		var form PostRoleUserForm
		err := c.ShouldBind(&form)
		if err == nil {
			res, resErr := sysRoleService.PostRoleUser(userId, form.Username, id)
			if resErr == nil {
				utils.OkDetailed(res, "success", c)
			} else {
				utils.FailWithMessage(resErr.Error(), c)
			}
		} else {
			utils.FailWithMessage(err.Error(), c)
		}
	}
}

// DeleteRoleUserByRoleIAndUserID 删除角色关联的用户
func DeleteRoleUserByRoleIDAndUserID(c *gin.Context) {
	if utils.CheckPermission(c, "DeleteRoleUserByRoleIDAndUserID") {
		id := utils.StrToUInt(c.Param("id"))
		userId := utils.StrToUInt(c.Param("userId"))
		resErr := sysRoleService.DeleteByRoleIdAndUserId(userId, id)
		if resErr == nil {
			utils.OkDetailed("删除成功！", "success", c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	}
}
