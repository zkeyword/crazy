package cms

import (
	"CRAZY/model/cms"
	"CRAZY/utils/db"
)

type CategoryRepository struct {
}

// Category 类型
type Category struct {
	ID     uint
	Title  string
	Status uint
	Index  uint
}

// NewUserRepository 实例化 DAO
func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

// Create 创建文章
func (r *CategoryRepository) Create(t *cms.Category) (*cms.Category, error) {
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
	return _db.Where("id = ?", id).Delete(Category{}).Error
}

// UpdateById 修改文章
func (r *CategoryRepository) UpdateById(id int64, t *cms.Category) (*cms.Category, error) {
	var ret = new(cms.Category)
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Model(&ret).Where("id=?", id).Updates(t).Error
	return ret, err
}

// GetById 获取文章
func (r *CategoryRepository) GetById(id int64) (*cms.Category, error) {
	var ret = &cms.Category{}
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.First(ret, "id = ?", id).Error
	return ret, err
}

// Get 获取文章
func (r *CategoryRepository) Get(page int, pageSize int, title string) ([]cms.Category, error) {
	var Category []cms.Category
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
	var Category []cms.Category
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
