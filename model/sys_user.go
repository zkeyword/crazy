package model

import (
	"time"
)

// User 用户表
type User struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Username  string `gorm:"unique;not null;varchar(50)"`
	Password  string `gorm:"not null;varchar(128)"`
	Status    int    `gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
