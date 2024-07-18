package sys

import (
	"time"
)

// SysPermission 权限表
type SysPermission struct {
	ID        uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name      string    `json:"name" gorm:"not null;unique;varchar(50)"`
	Key       string    `json:"key" gorm:"not null;unique;comment:'权限标识'"`
	PID       uint      `json:"pid" gorm:"default 0;comment:'父级'"`
	Status    int       `json:"status" gorm:"not null;comment:'用户状态: 1 - 正常、-1 - 禁用'"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
