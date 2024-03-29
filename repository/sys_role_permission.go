package repository

import (
	"CRAZY/model"
	"CRAZY/utils/db"
)

type RolePermissionRepository struct {
}

// RolePermission 类型
type RolePermission struct {
	ID             int64
	PermissionKeys string
	RoleId         int64
}

// NewRolePermissionRepository 实例化 DAO
func NewRolePermissionRepository() *RolePermissionRepository {
	return &RolePermissionRepository{}
}

func (r *RolePermissionRepository) Create(roleID uint, permissionKeys string) (*model.RolePermission, error) {
	var ret = new(model.RolePermission)
	ret.RoleID = roleID
	ret.PermissionKeys = permissionKeys
	err := db.GetMysql().Create(ret).Error
	return ret, err
}

func (r *RolePermissionRepository) DeleteById(id uint) error {
	if err := db.GetMysql().Where("id = ?", id).Delete(RolePermission{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *RolePermissionRepository) UpdateById(id uint, permissionKeys string) (*model.RolePermission, error) {
	var ret = new(model.RolePermission)
	data := &RolePermission{}
	data.PermissionKeys = permissionKeys
	err := db.GetMysql().Model(&ret).Where("role_id=?", id).Updates(data).Error
	return ret, err
}

func (r *RolePermissionRepository) Get(id int64) *RolePermission {
	ret := &RolePermission{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}
