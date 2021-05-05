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
	PermissionID uint
	RoleId       uint
}

// NewRolePermissionRepository 实例化 DAO
func NewRolePermissionRepository() *RolePermissionRepository {
	return &RolePermissionRepository{}
}

func (r *RolePermissionRepository) Create(RoleId uint, PermissionID uint) (*model.RolePermission, error) {
	var ret = new(model.RolePermission)
	ret.RoleID = RoleId
	ret.PermissionID = PermissionID
	err := db.GetMysql().Create(ret).Error
	return ret, err
}

func (r *RolePermissionRepository) DeleteById(id int64) error {
	if err := db.GetMysql().Where("id = ?", id).Delete(RolePermission{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *RolePermissionRepository) UpdateById(id int64, t *model.RolePermission) (*model.RolePermission, error) {
	var ret = new(model.RolePermission)
	err := db.GetMysql().Model(&ret).Where("id=?", id).Updates(t).Error
	return ret, err
}

func (r *RolePermissionRepository) Get(id int64) *RolePermission {
	ret := &RolePermission{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}
