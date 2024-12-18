package cms

import (
	"CRAZY/model/cms"
	"CRAZY/utils/db"
)

type CategoryRepository struct {
}

// NewUserRepository 实例化 DAO
func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

// Create 创建文章
func (r *CategoryRepository) Create(t *cms.CmsCategory) (*cms.CmsCategory, error) {
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Create(t).Error
	return t, err
}

// DeleteById 删除删除
func (r *CategoryRepository) DeleteById(id int64) error {
	_db, err := db.GetMysql()
	if err != nil {
		return err
	}
	return _db.Where("id = ?", id).Delete(cms.CmsCategory{}).Error
}

// UpdateById 修改文章
func (r *CategoryRepository) UpdateById(id int64, t *cms.CmsCategory) (*cms.CmsCategory, error) {
	var ret = new(cms.CmsCategory)
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Model(&ret).Where("id=?", id).Updates(t).Error
	return ret, err
}

// GetById 获取文章
func (r *CategoryRepository) GetById(id int64) (*cms.CmsCategory, error) {
	var ret = &cms.CmsCategory{}
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.First(ret, "id = ?", id).Error
	return ret, err
}

// Get 获取文章
func (r *CategoryRepository) Get(page int, pageSize int, title string) ([]cms.CmsCategory, error) {
	var Category []cms.CmsCategory
	var err error
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	if title != "" {
		err = _db.Order("`index` DESC, `updated_at` ASC").Where("title like ?", "%"+title+"%").Limit(pageSize).Offset((page - 1) * 10).Find(&Category).Error
	} else {
		err = _db.Order("`index` DESC, `updated_at` ASC").Limit(pageSize).Offset((page - 1) * 10).Find(&Category).Error
	}
	return Category, err
}

func (r *CategoryRepository) GetCategoryCount(title string) int64 {
	var Category []cms.CmsCategory
	var count int64
	_db, err := db.GetMysql()
	if err != nil {
		return 0
	}
	if title != "" {
		_db.Where("title like ?", "%"+title+"%").Find(&Category).Select("count(id)").Count(&count)
	} else {
		_db.Find(&Category).Select("count(id)").Count(&count)
	}
	return count
}
