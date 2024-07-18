package cms

import (
	"CRAZY/model/cms"
	"CRAZY/utils/db"
)

type BannerCategoryRepository struct {
}

// BannerCategory 类型
type BannerCategory struct {
	ID         uint
	BannerID   string
	CategoryID string
}

// NewUserRepository 实例化 DAO
func NewBannerCategoryRepository() *BannerCategoryRepository {
	return &BannerCategoryRepository{}
}

func (r *BannerCategoryRepository) Create(BannerID uint, categoryID uint) (*cms.CmsBannerCategory, error) {
	var ret = new(cms.CmsBannerCategory)
	ret.BannerID = BannerID
	ret.CategoryID = categoryID
	_db, err := db.GetMysql()
	if err != nil {
		return ret, err
	}
	err = _db.Create(ret).Error
	return ret, err
}

func (r *BannerCategoryRepository) DeleteByPostId(id int64) error {
	_db, err := db.GetMysql()
	if err != nil {
		return err
	}
	return _db.Where("banner_id = ?", id).Delete(BannerCategory{}).Error
}
