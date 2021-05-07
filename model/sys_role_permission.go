package model

import (
	"time"
)

// RolePermission 角色权限表
type RolePermission struct {
	ID             uint   `gorm:"primary_key;AUTO_INCREMENT"`
	RoleID         uint   `gorm:"not null"`
	PermissionKeys string `gorm:"not null;comment:'权限标识串'"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
