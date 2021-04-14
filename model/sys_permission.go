package model

import (
	"time"
)

// Permission 权限表
type Permission struct {
	ID        uint   `gorm:"primary_key;AUTO_INCREMENT"`
	Name      string `gorm:"not null;varchar(50)"`
	Action    int    `gorm:"-"`
	Status    int    `gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
