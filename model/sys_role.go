package model

import (
	"time"
)

// Role 角色表
type Role struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string `gorm:"unique;not null;varchar(50)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
