package cms

import (
	"time"
)

// CmsBannerCategory banner类别关联表
type CmsBannerCategory struct {
	ID         uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	BannerID   uint      `json:"bannerId" gorm:"not null"`
	CategoryID uint      `json:"categoryId" gorm:"not null;"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}
