package system

import (
	"time"
)

// SystemUserRole 用户角色表
type SystemUserRole struct {
	ID        uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserID    uint      `json:"userID" gorm:"not null"`
	RoleID    uint      `json:"roleID" gorm:"not null"`
	Username  string    `json:"username" gorm:"not null"` // Username 做冗余
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
