package model

import (
	"time"
)

// Role 角色表
type Role struct {
	ID        uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name      string    `json:"name" gorm:"unique;not null;varchar(50)"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
