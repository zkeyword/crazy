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
func (r *RoleRepository) Create(t *model.Role) (uint, error) {
	err := db.GetMysql().Create(t).Error
	return t.ID, err
}

// DeleteById 删除角色
func (r *RoleRepository) DeleteById(id int64) error {
	if err := db.GetMysql().Where("id = ?", id).Delete(Role{}).Error; err != nil {
		return err
	}

	return nil
}

// UpdateById 修改角色
func (r *RoleRepository) UpdateById(id int64, t *model.Role) (*model.Role, error) {
	var ret = new(model.Role)
	err := db.GetMysql().Model(&ret).Where("id=?", id).Updates(t).Error
	return ret, err
}

// Get 获取角色
func (r *RoleRepository) Get(id int64) *Role {
	ret := &Role{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}
