package sys

import (
	"CRAZY/model/sys"
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
func (r *PermissionRepository) Create(t *sys.SysPermission) (*sys.SysPermission, error) {
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Create(t).Error
	return t, err
}

// DeleteById 删除权限
func (r *PermissionRepository) DeleteById(id uint, key string) error {
	_db, err := db.GetMysql()
	if err != nil {
		return err
	}
	rolePermissionRepo := &RolePermissionRepository{}
	rolePermissionRepo.DeleteByKey(key)
	err = _db.Where("id = ?", id).Or("p_id = ?", id).Delete(Permission{}).Error
	return err
}

// UpdateById 修改权限
func (r *PermissionRepository) UpdateById(id uint, t *sys.SysPermission) (*sys.SysPermission, error) {
	var ret = new(sys.SysPermission)
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	rolePermissionRepo := &RolePermissionRepository{}
	rolePermissionRepo.DeleteByKey(ret.Key)
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
func (r *PermissionRepository) Get() []sys.SysPermission {
	var ret []sys.SysPermission
	_db, err := db.GetMysql()
	if err != nil {
		return nil
	}
	if err := _db.Find(&ret).Error; err != nil {
		return nil
	}

	return ret
}
