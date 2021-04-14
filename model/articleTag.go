package model

// ArticleTag 文章标签中间表
type ArticleTag struct {
	TagID     uint `gorm:"uint"`
	ArticleID uint `gorm:"uint"`
}
