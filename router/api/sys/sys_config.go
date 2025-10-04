package sys

import (
	"CRAZY/utils"
	"CRAZY/utils/xor"

	"github.com/gin-gonic/gin"
)

type ReturnGetUserPermissionKeys struct {
	PermissionKeys string `json:"permissions"`
	UserKey        string `json:"userKey"`
}

type PermItem struct {
	Key  string `json:"key"`  // 权限码（后台鉴权用）
	Name string `json:"name"` // 中文名称（前端展示用）
}

// GetUserPermissionKeys 获取用户权限
func GetUserPermissionKeys(c *gin.Context) {
	PermissionKeys, _ := c.Get("permissionKeys")
	PermissionKeysStr, _ := PermissionKeys.(string)
	userKey := utils.StringWithCharset(5)
	userPermissionKeys := &ReturnGetUserPermissionKeys{
		PermissionKeys: xor.XorEncryptDecrypt(PermissionKeysStr, userKey),
		UserKey:        userKey,
	}
	utils.OkDetailed(userPermissionKeys, "success", c)
}

// GetPermissionKeys 获取所有权限
func GetPermissionKeys(c *gin.Context) {
	list := []PermItem{
		// 用户
		{Key: "user:list", Name: "查看用户列表"},
		{Key: "user:view:byId", Name: "查看用户详情(byId)"},
		{Key: "user:view:byName", Name: "查看用户详情(byName)"},
		{Key: "user:add", Name: "新增用户"},
		{Key: "user:edit:byId", Name: "编辑用户信息(byId)"},
		{Key: "user:edit:status:byId", Name: "编辑用户状态(byId)"},
		{Key: "user:del", Name: "删除用户"},

		// 角色
		{Key: "role:list", Name: "查看角色列表"},
		{Key: "role:view:ById", Name: "查看角色详情(byId)"},
		{Key: "role:add", Name: "新增角色"},
		{Key: "role:edit:byId", Name: "编辑角色(byId)"},
		{Key: "role:del", Name: "删除角色"},

		// 权限
		{Key: "permission:tree:byId", Name: "查看权限树(byId)"},
		{Key: "permission:view:byId", Name: "查看权限详情(byId)"},
		{Key: "permission:add", Name: "添加权限"},
		{Key: "permission:edit:byId", Name: "编辑权限(byId)"},
		{Key: "permission:del", Name: "删除权限"},
		{Key: "permission:keys", Name: "获取所有权限列表"},

		// 关联
		{Key: "urp:view:byUserId", Name: "通过用户ID获取用户角色权限"},
		{Key: "rp:view:byRoleId", Name: "通过角色ID获取角色权限"},
		{Key: "rp:add:byRoleId", Name: "通过角色ID添加角色权限"},
		{Key: "ru:view:byUserId", Name: "通过用户ID获取用户角色"},
		{Key: "ru:view:byRoleId", Name: "通过角色ID获取角色用户"},
		{Key: "ru:edit:byRoleIDAndUserID", Name: "通过角色ID和用户ID编辑用户角色"},
		{Key: "ru:del:byRoleIDAndUserID", Name: "通过角色ID和用户ID删除用户角色"},

		// 数据字典
		{Key: "other:list", Name: "查看数据字典列表"},
		{Key: "other:view:byId", Name: "查看数据字典详情(byId)"},
		{Key: "other:add", Name: "添加数据字典"},
		{Key: "other:edit:byId", Name: "编辑数据字典(byId)"},
		{Key: "other:del", Name: "删除数据字典"},
	}

	utils.OkDetailed(list, "success", c)
}

// // GetPermissionKeys 获取所有权限
// func GetPermissionKeys(c *gin.Context) {
// 	permissionKeys := []string{
// 		// 用户
// 		"user:list",             // 查看用户列表
// 		"user:view:byId",        // 查看用户详情(byId)
// 		"user:view:byName",      // 查看用户详情(byName)
// 		"user:add",              // 新增用户
// 		"user:edit:byId",        // 编辑用户信息(byId)
// 		"user:edit:status:byId", // 编辑用户status(byId)
// 		"user:del",              // 删除用户(byId)

// 		// 角色
// 		"role:list",      // 查看角色列表
// 		"role:view:ById", // 查看角色详情(byId)
// 		"role:add",       // 新增角色
// 		"role:edit:byId", // 编辑角色(byId)
// 		"role:del",       // 删除角色(byId)

// 		// 权限
// 		// "GetPermissionById",
// 		// "GetPermissionTreeById",
// 		// "PostPermission",
// 		// "DelPermissionById",
// 		// "PutPermissionById",
// 		// "GetPermissionKeys",

// 		"permission:tree:byId", // 查看权限树(byId)
// 		"permission:view:byId", // 查看权限详情(byId)
// 		"permission:add",       // 添加权限
// 		"permission:edit:byId", // 编辑权限(byId)
// 		"permission:del",       // 删除权限(byId)
// 		"permission:keys",      // 获取所有权限列表

// 		// 角色、用户、权限关联
// 		// "GetUserRolePermissionByUserId",
// 		// "GetRolePermissionByRoleID",
// 		// "PostRolePermissionByRoleID",
// 		// "GetRoleUserByUserID",
// 		// "GetRoleUserByRoleID",
// 		// "PostRoleUserByRoleIDAndUserID",
// 		// "DeleteRoleUserByRoleIDAndUserID",
// 		"urp:view:byUserId",         // 通过用户ID获取用户角色权限
// 		"rp:view:byRoleId",          // 通过角色ID获取角色权限
// 		"rp:add:byRoleId",           // 通过角色ID添加角色权限
// 		"ru:view:byUserId",          // 通过用户ID获取用户角色
// 		"ru:view:byRoleId",          // 通过角色ID获取角色用户
// 		"ru:edit:byRoleIDAndUserID", // 通过角色ID和用户ID编辑用户角色
// 		"ru:del:byRoleIDAndUserID",  // 通过角色ID和用户ID删除用户角色

// 		// 其他设置
// 		"other:list",      // 查看其他列表
// 		"other:view:byId", // 查看其他详情(byId)
// 		"other:add",       // 添加其他
// 		"other:edit:byId", // 编辑其他(byId)
// 		"other:del",       // 删除其他(byId)
// 		// "GetOther",
// 		// "GetOtherById",
// 		// "PostOther",
// 		// "DelOtherById",
// 		// "PutOtherById",
// 		// const (
// 		// 	PermOtherGet   Permission = "other:get"
// 		// 	PermOtherByID  Permission = "other:byId"
// 		// 	PermOtherPost  Permission = "other:post"
// 		// 	PermOtherPut   Permission = "other:put"
// 		// 	PermOtherDel   Permission = "other:del"
// 		// 	PermKeys       Permission = "perm:keys"
// 		// )

// 	}
// 	utils.OkDetailed(permissionKeys, "success", c)
// }
