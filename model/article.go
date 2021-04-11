package model

import (
	"github.com/jinzhu/gorm"
)

// Article 文章表
type Article struct {
	gorm.Model
	Title    string `gorm:"size:128;not null;"`      // 文章标题
	Content  string `gorm:"type:longtext;not null;"` // 文章内容
	AuthorID int    `gorm:"int"`                     // 作者
}
