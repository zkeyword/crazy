package model

import (
	"time"
)

// RolePermission 角色权限表
type RolePermission struct {
	ID           uint `gorm:"primary_key;AUTO_INCREMENT"`
	PermissionId int  `gorm:"-"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
