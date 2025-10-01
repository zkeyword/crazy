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
	permissionKeys := []string{
		// 用户
		"GetUser",
		"GetUserByUsername",
		"GetUserById",
		"PostUser",
		"PutUserById",
		"PutUserStatusById",
		// 角色
		"GetRole",
		"GetRoleById",
		"PostRole",
		"DelRoleById",
		"PutRoleById",
		// 权限
		"GetPermissionById",
		"GetPermissionTreeById",
		"PostPermission",
		"DelPermissionById",
		"PutPermissionById",
		// 角色、用户、权限关联
		"GetRolePermissionByRoleID",
		"PostRolePermissionByRoleID",
		"GetRoleUserByUserID",
		"GetUserRolePermissionByUserId",
		"GetRoleUserByRoleID",
		"PostRoleUserByRoleIDAndUserID",
		"DeleteRoleUserByRoleIDAndUserID",
		// 其他设置
		"GetOther",
		"GetOtherById",
		"PostOther",
		"DelOtherById",
		"PutOtherById",
		//
		"GetPermissionKeys",
	}
	utils.OkDetailed(permissionKeys, "success", c)
}
