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
	err := db.GetMysql().Create(t).Error
	return t, err
}

// DeleteById 删除权限
func (r *PermissionRepository) DeleteById(id uint) error {
	if err := db.GetMysql().Where("id = ?", id).Delete(Permission{}).Error; err != nil {
		return err
	}

	return nil
}

// UpdateById 修改权限
func (r *PermissionRepository) UpdateById(id uint, t *model.Permission) (*model.Permission, error) {
	var ret = new(model.Permission)
	err := db.GetMysql().Model(&ret).Where("id=?", id).Updates(t).Error
	return ret, err
}

// GetById 获取权限
func (r *PermissionRepository) GetById(id uint) *Permission {
	ret := &Permission{}
	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}

// Get 获取权限列表
func (r *PermissionRepository) Get() []model.Permission {
	var ret []model.Permission

	if err := db.GetMysql().Find(&ret).Error; err != nil {
		return nil
	}

	return ret
}
