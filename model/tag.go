package model

import (
	"github.com/jinzhu/gorm"
)

// Tag 标签表
type Tag struct {
	gorm.Model
	Title string `gorm:"size:128;not null;"`
}
