package model

import (
	"time"
)

// UserRole 用户角色表
type UserRole struct {
	ID        uint `gorm:"primary_key;AUTO_INCREMENT"`
	UserID    uint `gorm:"not null"`
	RoleID    uint `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
