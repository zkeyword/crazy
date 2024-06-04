package repository

import (
	"CRAZY/model"
	"CRAZY/utils/db"
)

type RolePermissionRepository struct {
}

// RolePermission 类型
type RolePermission struct {
	ID             uint   `json:"id"`
	PermissionKeys string `json:"permissionKeys"`
	RoleId         uint   `json:"roleId"`
}

// NewRolePermissionRepository 实例化 DAO
func NewRolePermissionRepository() *RolePermissionRepository {
	return &RolePermissionRepository{}
}

func (r *RolePermissionRepository) Create(roleID uint, permissionKeys string) (*model.RolePermission, error) {
	var ret = new(model.RolePermission)
	ret.RoleID = roleID
	ret.PermissionKeys = permissionKeys
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Create(ret).Error
	return ret, err
}

func (r *RolePermissionRepository) DeleteById(id uint) error {
	_db, err := db.GetMysql()
	if err != nil {
		return nil
	}
	return _db.Where("id = ?", id).Delete(RolePermission{}).Error
}

func (r *RolePermissionRepository) UpdateByRoleId(id uint, permissionKeys string) (*model.RolePermission, error) {
	var ret = new(model.RolePermission)
	data := &RolePermission{}
	data.PermissionKeys = permissionKeys
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Model(&ret).Where("role_id=?", id).Updates(data).Error
	return ret, err
}

func (r *RolePermissionRepository) GetByRoleID(id uint) *RolePermission {
	ret := &RolePermission{}
	_db, err := db.GetMysql()
	if err != nil {
		return nil
	}
	if err := _db.First(ret, "role_id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}
