package sys

import (
	"CRAZY/model/sys"
	sysUserService "CRAZY/services/sys/sys_user"
	"CRAZY/utils"
	"CRAZY/utils/xor"
	"html"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserForm struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
	RealName string `form:"realName" binding:"required"`
	Status   *int   `form:"status" binding:"required,gte=0"`
	RoleIDs  string `form:"roleIds"`
}

// PostUser 新增用户
func PostUser(c *gin.Context) {
	if utils.CheckPermission(c, "user:add") {
		var form UserForm
		err := c.ShouldBind(&form)
		if form.RoleIDs == "" {
			// TODO: https://www.cnblogs.com/xinliangcoder/p/11234017.html 自定义验证器
			utils.FailWithMessage("Key: 'UserForm.RoleIDs' Error:Field validation for 'roleIds' failed on the 'lt' 0", c)
			return
		}
		if err == nil && form.Status != nil {
			Model := &sys.SysUser{
				Username: html.EscapeString(form.Username),
				Password: xor.Enc(form.Password),
				RealName: form.RealName,
				Status:   *form.Status,
			}
			res, resErr := sysUserService.Create(Model, form.RoleIDs)
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

// DelUserById 删除用户
func DelUserById(c *gin.Context) {
	if utils.CheckPermission(c, "user:del") {
		id := utils.StrToUInt(c.Param("id"))
		resErr := sysUserService.DeleteById(id)
		if resErr == nil {
			utils.Ok(c)
		} else {
			utils.FailWithMessage(resErr.Error(), c)
		}
	}
}

// PutUserById 修改用户
func PutUserById(c *gin.Context) {
	if utils.CheckPermission(c, "user:edit:byId") {
		id := utils.StrToUInt(c.Param("id"))
		var form UserForm
		err := c.ShouldBind(&form)
		if err == nil && form.Status != nil {
			userRes, _ := sysUserService.GetById(id)
			var Password = userRes.Password
			if userRes.Password != form.Password {
				Password = xor.Enc(form.Password)
			}
			LoginStatus := 0 // 为 0 gorm 忽略， Status 为禁用时，设置 LoginStatus 是登出
			if *form.Status == -1 {
				LoginStatus = -1
			}
			Model := &sys.SysUser{
				Username:    form.Username,
				Password:    Password,
				RealName:    form.RealName,
				Status:      *form.Status,
				LoginStatus: LoginStatus,
			}
			res, resErr := sysUserService.PutUserById(id, Model, form.RoleIDs)
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

type UserDisableForm struct {
	Status *int `form:"status" binding:"required"`
}

// PutUserStatusById 修改用户状态
func PutUserStatusById(c *gin.Context) {
	if utils.CheckPermission(c, "user:edit:status:byId") {
		id := utils.StrToUInt(c.Param("id"))
		var form UserDisableForm
		err := c.ShouldBind(&form)
		if err == nil && form.Status != nil {
			LoginStatus := 0 // 为 0 gorm 忽略， Status 为禁用时，设置 LoginStatus 是登出
			if *form.Status == -1 {
				LoginStatus = -1
			}
			Model := &sys.SysUser{
				Status:      *form.Status,
				LoginStatus: LoginStatus,
			}
			res, resErr := sysUserService.PutUserById(id, Model, "")
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

// GetUser 获取用户列表
func GetUser(c *gin.Context) {
	if utils.CheckPermission(c, "user:list") {
		page, _ := strconv.Atoi(c.Query("page"))
		pageSize, _ := strconv.Atoi(c.Query("pageSize"))
		username := html.EscapeString(c.Query("username"))
		res, _ := sysUserService.Get(page, pageSize, username)
		utils.OkDetailed(res, "success", c)
	}
}

// GetUserById 获取用户
func GetUserById(c *gin.Context) {
	if utils.CheckPermission(c, "user:view:byId") {
		id := utils.StrToUInt(c.Param("id"))
		res, _ := sysUserService.GetById(id)
		// userDetailRes := sysUserService.GetUserRolePermissionByUserId(res.ID)
		utils.OkDetailed(res, "success", c)
	}
}

// GetUserByUsername 获取用户
func GetUserByUsername(c *gin.Context) {
	if utils.CheckPermission(c, "user:view:byName") {
		username := c.Param("username")
		res, _ := sysUserService.GetByUserName(html.EscapeString(username))
		utils.OkDetailed(res, "success", c)
	}
}

// GetUserRolePermissionByUserId 获取用户角色权限
func GetUserRolePermissionByUserId(c *gin.Context) {
	if utils.CheckPermission(c, "GetUserRolePermissionByUserId") {
		id := utils.StrToUInt(c.Param("id"))
		res := sysUserService.GetUserRolePermissionByUserId(id)
		utils.OkDetailed(res, "success", c)
	}
}
