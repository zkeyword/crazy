package shop

import (
	"time"
)

// ShopUser 用户表
type ShopUser struct {
	ID        uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Username  string    `json:"username" gorm:"unique;not null;varchar(50)"`
	Password  string    `json:"password" gorm:"not null;varchar(128)"`
	Status    int       `json:"status" gorm:"not null;comment:'用户状态: 1 - 正常、-1 - 禁用'"`
	Level     int       `json:"level" gorm:"not null;comment:'用户等级'"`
	ParentID  uint      `json:"parentID" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
