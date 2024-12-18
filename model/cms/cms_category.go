package cms

import (
	"time"
)

// category 分类表
type CmsCategory struct {
	ID        uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Status    int       `json:"status" gorm:"not null;comment:'用户状态: 1 - 正常、-1 - 禁用'"`
	Index     uint      `json:"index" gorm:"not null"`
	Title     string    `json:"title" gorm:"type:varchar(255);unique;not null;"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
