package model

import (
	"time"
)

// Permission 权限表
type Permission struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string `gorm:"not null;unique;varchar(50)"`
	Key       string `gorm:"not null;unique;comment:'权限标识'"`
	Status    int    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
