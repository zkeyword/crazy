package repository

import (
	"CRAZY/model"
	"CRAZY/utils/db"
)

type UserAddressRepository struct {
}

// User 类型
type UserAddress struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// NewUserAddressRepository 实例化 DAO
func NewUserAddressRepository() *UserAddressRepository {
	return &UserAddressRepository{}
}

func (r *UserAddressRepository) Create(t *model.User) (*model.User, error) {
	err := db.GetMysql().Create(t).Error
	return t, err
}

// DeleteById 删除用户
func (r *UserAddressRepository) DeleteById(id uint) error {
	if err := db.GetMysql().Where("id = ?", id).Delete(User{}).Error; err != nil {
		return err
	}

	return nil
}

// UpdateById 修改用户
func (r *UserAddressRepository) UpdateById(id uint, t *model.User) (*model.User, error) {
	var ret = new(model.User)
	err := db.GetMysql().Model(&ret).Where("id=?", id).Updates(t).Error
	return ret, err
}

// Get 获取用户列表
func (r *UserAddressRepository) Get(page int, pageSize int, username string) ([]model.User, error) {
	var users []model.User
	var err error
	if pageSize < 1 {
		pageSize = 10
	}
	if username != "" {
		err = db.GetMysql().Where("username like ?", "%"+username+"%").Limit(pageSize).Offset((page - 1) * pageSize).Find(&users).Error
	} else {
		err = db.GetMysql().Limit(pageSize).Offset((page - 1) * pageSize).Find(&users).Error
	}
	return users, err
}
