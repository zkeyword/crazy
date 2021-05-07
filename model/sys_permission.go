package model

import (
	"time"
)

// Permission 权限表
type Permission struct {
	ID        uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Name      string    `json:"name" gorm:"not null;unique;varchar(50)"`
	Key       string    `json:"key" gorm:"not null;unique;comment:'权限标识'"`
	Status    int       `json:"status" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
