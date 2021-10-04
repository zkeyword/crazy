package model

import (
	"time"
)

// Log 日志表
type Log struct {
	ID        uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserID    string    `json:"userId" gorm:"not null;"`
	Type      uint      `json:"type" gorm:"unit;not null"`
	Value     string    `json:"value" gorm:"type:varchar(255);not null;"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
