package model

import (
	"time"
)

// UserRole 用户角色表
type UserRole struct {
	ID        uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserID    uint      `json:"userId" gorm:"not null"`
	RoleID    uint      `json:"roleId" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
