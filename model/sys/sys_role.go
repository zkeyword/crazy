package sys

import (
	"time"
)

// SysRole 角色表
type SysRole struct {
	ID        uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name      string    `json:"name" gorm:"unique;not null;varchar(50)"`
	Desc      string    `json:"desc" gorm:"type:varchar(255);comment:'角色描述'"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
