package model

import (
	"time"
)

// RolePermission 角色权限表
type RolePermission struct {
	ID           uint `gorm:"primary_key;AUTO_INCREMENT"`
	RoleID       uint `gorm:"not null"`
	PermissionID uint `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
