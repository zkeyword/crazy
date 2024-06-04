package repository

import (
	"CRAZY/model"
	"CRAZY/utils/db"
	"time"
)

type PermissionRepository struct {
}

type Permission struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Key       string    `json:"key"`
	Status    int       `json:"status"`
	PID       int       `json:"pid"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

func NewPermissionRepository() *PermissionRepository {
	return &PermissionRepository{}
}

// Create 创建权限
func (r *PermissionRepository) Create(t *model.Permission) (*model.Permission, error) {
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Create(t).Error
	return t, err
}

// DeleteById 删除权限
func (r *PermissionRepository) DeleteById(id uint) error {
	_db, err := db.GetMysql()
	if err != nil {
		return err
	}
	err = _db.Where("id = ?", id).Delete(Permission{}).Error
	return err
}

// UpdateById 修改权限
func (r *PermissionRepository) UpdateById(id uint, t *model.Permission) (*model.Permission, error) {
	var ret = new(model.Permission)
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Model(&ret).Where("id=?", id).Updates(t).Error
	if err == nil {
		ret.ID = id
	}
	return ret, err
}

// GetById 获取权限
func (r *PermissionRepository) GetById(id uint) *Permission {
	ret := &Permission{}
	_db, err := db.GetMysql()
	if err != nil {
		return nil
	}
	err = _db.First(ret, "id = ?", id).Error
	if err != nil {
		return nil
	}
	return ret
}

// Get 获取权限列表
func (r *PermissionRepository) Get() []model.Permission {
	var ret []model.Permission
	_db, err := db.GetMysql()
	if err != nil {
		return nil
	}
	if err := _db.Find(&ret).Error; err != nil {
		return nil
	}

	return ret
}
