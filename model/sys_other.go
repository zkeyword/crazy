package model

import (
	"time"
)

// Other
type Other struct {
	ID        uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Key       string    `json:"key" gorm:"type:varchar(255);unique;not null;"`
	Value     string    `json:"value" gorm:"type:varchar(255);not null;"`
	Type      uint      `json:"type" gorm:"unit;not null"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
