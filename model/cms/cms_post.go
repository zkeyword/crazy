package cms

import (
	"time"
)

type CmsPost struct {
	ID          uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Title       string    `json:"title" gorm:"type:varchar(255);unique;not null;"`
	Desc        string    `json:"desc" gorm:"type:varchar(255);unique;not null;"`
	Thumbnail   string    `json:"thumbnail" gorm:"type:varchar(255);not null"`
	Content     string    `json:"content" gorm:"type:text;not null;text"`
	Status      int       `json:"status" gorm:"not null;comment:'用户状态: 1 - 正常、-1 - 禁用'"`
	Index       uint      `json:"index" gorm:"unit;not null"`
	CategoryIds string    `json:"categoryIds" gorm:"type:varchar(50);not null"`
	NewsRank    int       `json:"news_rank"`
	PublishAt   int       `json:"publishAt"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
