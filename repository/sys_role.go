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
func (r *RoleRepository) GetById(id uint) *Role {
	ret := &Role{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
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
	if name != "" {
		err = db.GetMysql().Where("name like ?", "%"+name+"%").Limit(pageSize).Offset((page - 1) * pageSize).Find(&roles).Error
	} else {
		err = db.GetMysql().Limit(pageSize).Offset((page - 1) * pageSize).Find(&roles).Error
	}
	return roles, err
}

func (r *RoleRepository) GetRoleCount(name string) int {
	var roles []model.Role
	var count int
	if name != "" {
		db.GetMysql().Where("name like ?", "%"+name+"%").Find(&roles).Select("count(id)").Count(&count)
	} else {
		db.GetMysql().Find(&roles).Select("count(id)").Count(&count)
	}
	return count
}
