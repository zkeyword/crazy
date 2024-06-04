package repository

import (
	"CRAZY/model"
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
func (r *RoleRepository) Create(t *model.Role) (*model.Role, error) {
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
	if err := _db.Where("id = ?", id).Delete(Role{}).Error; err != nil {
		return err
	}

	return nil
}

// UpdateById 修改角色
func (r *RoleRepository) UpdateById(id uint, t *model.Role) (*model.Role, error) {
	var ret = new(model.Role)
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
func (r *RoleRepository) Get(page int, pageSize int, name string) ([]model.Role, error) {
	var roles []model.Role
	var err error
	if pageSize < 1 {
		pageSize = 10
	}
	if page < 1 {
		page = 1
	}
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	if name != "" {
		err = _db.Where("name like ?", "%"+name+"%").Limit(pageSize).Offset((page - 1) * pageSize).Find(&roles).Error
	} else {
		err = _db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&roles).Error
	}
	return roles, err
}

func (r *RoleRepository) GetRoleCount(name string) int64 {
	var roles []model.Role
	var count int64
	_db, err := db.GetMysql()
	if err != nil {
		return 0
	}
	if name != "" {
		_db.Where("name like ?", "%"+name+"%").Find(&roles).Select("count(id)").Count(&count)
	} else {
		_db.Find(&roles).Select("count(id)").Count(&count)
	}
	return count
}
