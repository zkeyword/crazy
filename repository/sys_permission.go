package repository

import (
	"CRAZY/model"
	"CRAZY/utils/db"
	"time"
)

type PermissionRepository struct {
}

type PermissionList struct {
	Data     []model.Permission
	Total    int
	PageSize int
	Page     int
}

type Permission struct {
	ID        uint
	Title     string
	Content   string
	AuthorID  uint
	UpdatedAt time.Time
	TagID     int
	TagName   string
}

func NewPermissionRepository() *PermissionRepository {
	return &PermissionRepository{}
}

// Create 创建权限
func (r *PermissionRepository) Create(t *model.Permission) (uint, error) {
	err := db.GetMysql().Create(t).Error
	return t.ID, err
}

// DeleteById 删除权限
func (r *PermissionRepository) DeleteById(id int64) error {
	if err := db.GetMysql().Where("id = ?", id).Delete(Permission{}).Error; err != nil {
		return err
	}

	return nil
}

// UpdateById 修改权限
func (r *PermissionRepository) UpdateById(id int64, t *model.Permission) (*model.Permission, error) {
	var ret = new(model.Permission)
	err := db.GetMysql().Model(&ret).Where("id=?", id).Updates(t).Error
	return ret, err
}

// Get 获取权限
func (r *PermissionRepository) Get(id int64) *Permission {
	ret := &Permission{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}
