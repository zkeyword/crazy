package cms

import (
	"time"
)

// Banner表
type CmsBanner struct {
	ID          uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Status      int       `json:"status" gorm:"not null;comment:'用户状态: 1 - 正常、-1 - 禁用'"`
	Index       uint      `json:"index" gorm:"not null"`
	Type        uint      `json:"type" gorm:"not null"`
	Title       string    `json:"title" gorm:"type:varchar(255);unique;not null;"`
	Desc        string    `json:"desc" gorm:"type:varchar(255);unique;not null;"`
	Image       string    `json:"image" gorm:"type:varchar(255);not null"`
	Link        string    `json:"link" gorm:"type:varchar(255);not null"`
	CategoryIds string    `json:"categoryIds" gorm:"type:varchar(50);not null"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
