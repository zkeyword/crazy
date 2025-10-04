package sys

import (
	"CRAZY/model/sys"
	"CRAZY/utils/db"
	"strings"
)

type RolePermissionRepository struct {
}

// RolePermission 类型
type SysRolePermission struct {
	ID             uint   `json:"id"`
	PermissionKeys string `json:"permissionKeys"`
	RoleId         uint   `json:"roleId"`
}

// NewRolePermissionRepository 实例化 DAO
func NewRolePermissionRepository() *RolePermissionRepository {
	return &RolePermissionRepository{}
}

func (r *RolePermissionRepository) Create(roleID uint, permissionKeys string) ([]*sys.SysRolePermission, error) {
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	keys := strings.Split(permissionKeys, ",")

	var rolePermissions []*sys.SysRolePermission
	for _, key := range keys {
		rolePermission := &sys.SysRolePermission{
			RoleID:         roleID,
			PermissionKeys: key,
		}
		rolePermissions = append(rolePermissions, rolePermission)
	}

	err = _db.Create(rolePermissions).Error
	return rolePermissions, err
}

func (r *RolePermissionRepository) DeleteByRoleId(id uint) error {
	_db, err := db.GetMysql()
	if err != nil {
		return nil
	}
	return _db.Where("role_id = ?", id).Delete(SysRolePermission{}).Error
}

func (r *RolePermissionRepository) DeleteByKey(key string) error {
	_db, err := db.GetMysql()
	if err != nil {
		return nil
	}
	return _db.Where("permission_keys = ?", key).Delete(SysRolePermission{}).Error
}

func (r *RolePermissionRepository) UpdateByRoleId(id uint, permissionKeys string) ([]*sys.SysRolePermission, error) {
	r.DeleteByRoleId(id)
	ret, err := r.Create(id, permissionKeys)
	return ret, err
}

func (r *RolePermissionRepository) GetByRoleID(id uint) *SysRolePermission {
	var ret = &SysRolePermission{}
	rets := &[]SysRolePermission{}
	_db, err := db.GetMysql()
	if err != nil {
		return nil
	}
	if err := _db.Find(rets, "role_id = ?", id).Error; err != nil {
		return nil
	}
	var keys []string
	for _, val := range *rets {
		ret.ID = val.ID
		keys = append(keys, val.PermissionKeys)
	}
	ret.PermissionKeys = strings.Join(keys, ",")
	return ret
}
