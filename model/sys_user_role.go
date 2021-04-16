package model

import (
	"time"
)

// UserRole 用户角色表
type UserRole struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT"`
	UserId    uint `gorm:"-"`
	RoleId    uint `gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
