package model

import (
	"time"
)

// User 用户表
type User struct {
	ID        uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Username  string    `json:"username" gorm:"unique;not null;varchar(50)"`
	Password  string    `json:"password" gorm:"not null;varchar(128)"`
	Status    int       `json:"status" gorm:"not null"`
	Level     int       `json:"level" gorm:"not null"`
	ParentID  uint      `json:"parentID" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
