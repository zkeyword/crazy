package cms

import (
	"CRAZY/model/cms"
	"CRAZY/utils/db"
)

type BannerRepository struct {
}

// Banner 类型
type Banner struct {
	ID        uint
	Title     string
	Thumbnail string
	Content   string
	Type      uint
	Index     uint
}

// NewUserRepository 实例化 DAO
func NewBannerRepository() *BannerRepository {
	return &BannerRepository{}
}

// Create 创建文章
func (r *BannerRepository) Create(t *cms.Banner) (*cms.Banner, error) {
	_db, err := db.GetMysql()
	if err != nil {
		return t, err
	}
	err = _db.Create(t).Error
	return t, err
}

// DeleteById 删除删除
func (r *BannerRepository) DeleteById(id int64) error {
	_db, err := db.GetMysql()
	if err != nil {
		return err
	}

	return _db.Where("id = ?", id).Delete(Banner{}).Error
}

// UpdateById 修改文章
func (r *BannerRepository) UpdateById(id int64, t *cms.Banner) (*cms.Banner, error) {
	var ret = new(cms.Banner)
	_db, err := db.GetMysql()
	if err != nil {
		return ret, err
	}
	err = _db.Model(&ret).Where("id=?", id).Updates(t).Error
	return ret, err
}

// GetById 获取文章
func (r *BannerRepository) GetById(id int64) (*cms.Banner, error) {
	var ret = &cms.Banner{}
	_db, err := db.GetMysql()
	if err != nil {
		return ret, err
	}
	err = _db.First(ret, "id = ?", id).Error
	return ret, err
}

// Get 获取文章
func (r *BannerRepository) Get(page int, pageSize int, title string) ([]cms.Banner, error) {
	var Banner []cms.Banner
	var err error
	_db, err := db.GetMysql()
	if err != nil {
		return Banner, err
	}
	if title != "" {
		err = _db.Order("`index` DESC, `updated_at` ASC").Where("title like ?", "%"+title+"%").Limit(pageSize).Offset((page - 1) * 10).Find(&Banner).Error
	} else {
		err = _db.Order("`index` DESC, `updated_at` ASC").Limit(pageSize).Offset((page - 1) * 10).Find(&Banner).Error
	}
	return Banner, err
}

func (r *BannerRepository) GetBannerCount(title string) int64 {
	var users []cms.Banner
	var count int64
	_db, err := db.GetMysql()
	if err != nil {
		return 0
	}
	if title != "" {
		_db.Where("title like ?", "%"+title+"%").Find(&users).Select("count(id)").Count(&count)
	} else {
		_db.Find(&users).Select("count(id)").Count(&count)
	}
	return count
}
