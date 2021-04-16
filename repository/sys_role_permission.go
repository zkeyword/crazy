package repository

import (
	"CRAZY/model"
	"CRAZY/utils/db"
)

type RolePermissionRepository struct {
}

// RolePermission 类型
type RolePermission struct {
	ID           uint
	PermissionId string
	RoleId       string
}

// NewRolePermissionRepository 实例化 DAO
func NewRolePermissionRepository() *RolePermissionRepository {
	return &RolePermissionRepository{}
}

// Create 创建用户
func (r *RolePermissionRepository) Create(t *model.RolePermission) (uint, error) {
	err := db.GetMysql().Create(t).Error
	return t.ID, err
}

// DeleteById 删除用户
func (r *RolePermissionRepository) DeleteById(id int64) error {
	if err := db.GetMysql().Where("id = ?", id).Delete(RolePermission{}).Error; err != nil {
		return err
	}

	return nil
}

// UpdateById 修改用户
func (r *RolePermissionRepository) UpdateById(id int64, t *model.RolePermission) (*model.RolePermission, error) {
	var ret = new(model.RolePermission)
	err := db.GetMysql().Model(&ret).Where("id=?", id).Updates(t).Error
	return ret, err
}

// Get 获取用户
func (r *RolePermissionRepository) Get(id int64) *RolePermission {
	ret := &RolePermission{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}
