package repository

import (
	"CRAZY/model"
	"CRAZY/utils/db"
)

type RoleRepository struct {
}

type Role struct {
	ID   uint
	Name string
}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{}
}

// Create 创建角色
func (r *RoleRepository) Create(t *model.Role) (*model.Role, error) {
	err := db.GetMysql().Create(t).Error
	return t, err
}

// DeleteById 删除角色
func (r *RoleRepository) DeleteById(id uint) error {
	if err := db.GetMysql().Where("id = ?", id).Delete(Role{}).Error; err != nil {
		return err
	}

	return nil
}

// UpdateById 修改角色
func (r *RoleRepository) UpdateById(id uint, t *model.Role) (*model.Role, error) {
	var ret = new(model.Role)
	err := db.GetMysql().Model(&ret).Where("id=?", id).Updates(t).Error
	t.ID = id
	return t, err
}

// Get 获取角色
func (r *RoleRepository) Get(id uint) *Role {
	ret := &Role{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}
