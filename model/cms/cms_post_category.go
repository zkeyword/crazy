package cms

import (
	"time"
)

// PostCategory 文章类别关联表
type PostCategory struct {
	ID         uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	PostID     uint      `json:"postId" gorm:"not null"`
	CategoryID uint      `json:"categoryId" gorm:"not null;"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
