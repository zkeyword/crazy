package sys

import (
	"time"
)

// SysUser 用户表
type SysUser struct {
	ID          uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Username    string    `json:"username" gorm:"unique;not null;varchar(50)"`
	RealName    string    `json:"realName" gorm:"unique;not null;varchar(50)"`
	Password    string    `json:"password" gorm:"not null;varchar(128)"`
	Status      int       `json:"status" gorm:"not null;comment:'用户状态: 1 - 正常、-1 - 禁用'"`
	LoginStatus int       `json:"loginStatus" gorm:"not null;comment:'用户状态: 1 - 正常、-1 - 禁用'"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
