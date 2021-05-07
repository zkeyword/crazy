package model

import (
	"time"
)

// RolePermission 角色权限表
type RolePermission struct {
	ID             uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	RoleID         uint      `json:"roleID" gorm:"not null"`
	PermissionKeys string    `json:"permissionKeys" gorm:"not null;comment:'权限标识串'"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}
