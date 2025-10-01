package sys

import (
	"CRAZY/model/sys"
	"CRAZY/utils/db"
)

type RoleRepository struct {
}

type Role struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Desc string `json:"desc"`
}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{}
}

// Create 创建角色
func (r *RoleRepository) Create(t *sys.SysRole) (*sys.SysRole, error) {
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Create(t).Error
	return t, err
}

// DeleteById 删除角色
func (r *RoleRepository) DeleteById(id uint) error {
	_db, err := db.GetMysql()
	if err != nil {
		return err
	}
	return _db.Where("id = ?", id).Delete(Role{}).Error
}

// UpdateById 修改角色
func (r *RoleRepository) UpdateById(id uint, t *sys.SysRole) (*sys.SysRole, error) {
	var ret = new(sys.SysRole)
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Model(&ret).Where("id=?", id).Updates(t).Error
	if err == nil {
		t.ID = id
	}
	return t, err
}

// Get 获取角色
func (r *RoleRepository) GetById(id uint) *Role {
	ret := &Role{}
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

// Get 获取角色列表
func (r *RoleRepository) Get(page int, pageSize int, name string) ([]sys.SysRole, error) {
	var roles []sys.SysRole

	// 默认值处理
	if pageSize < 1 {
		pageSize = 10
	}
	if page < 1 {
		page = 1
	}

	// 获取数据库连接
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}

	// 构建查询
	query := _db.Model(&sys.SysRole{})
	if name != "" {
		query = query.Where("name LIKE ?", "%"+name+"%")
	}

	// 分页查询
	err = query.Order("id ASC").Limit(pageSize).Offset((page - 1) * pageSize).Find(&roles).Error
	if err != nil {
		return nil, err
	}

	return roles, nil
}

func (r *RoleRepository) GetRoleCount(name string) int64 {
	var roles []sys.SysRole
	var count int64
	_db, err := db.GetMysql()
	if err != nil {
		return 0
	}
	if name != "" {
		_db.Where("name like ?", "%"+name+"%").Model(&roles).Select("count(id)").Count(&count)
	} else {
		_db.Model(&roles).Select("count(id)").Count(&count)
	}
	return count
}
