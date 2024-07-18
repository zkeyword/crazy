package system

import (
	"time"
)

// SystemRole 角色表
type SystemRole struct {
	ID        uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name      string    `json:"name" gorm:"unique;not null;varchar(50)"`
	Desc      string    `json:"desc" gorm:"type:varchar(255);comment:'角色描述'"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
