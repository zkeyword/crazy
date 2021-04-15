package repository

import (
	"CRAZY/model"
	"CRAZY/utils/db"
)

type UserRepository struct {
}

// User 类型
type User struct {
	ID       uint
	Username string
	Password string
}

// NewUserRepository 实例化 DAO
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Create 创建用户
func (r *UserRepository) Create(t *model.User) (uint, error) {
	err := db.GetMysql().Create(t).Error
	return t.ID, err
}

// DeleteById 删除用户
func (r *UserRepository) DeleteById(id int64) error {
	if err := db.GetMysql().Where("id = ?", id).Delete(User{}).Error; err != nil {
		return err
	}

	return nil
}

// UpdateById 修改用户
func (r *UserRepository) UpdateById(id int64, t *model.User) (*model.User, error) {
	var ret = new(model.User)
	err := db.GetMysql().Model(&ret).Where("id=?", id).Updates(t).Error
	return ret, err
}

// Get 获取用户
func (r *UserRepository) Get(id int64) *User {
	ret := &User{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}
