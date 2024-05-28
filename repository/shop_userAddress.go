package repository

import (
	"CRAZY/model"
	"CRAZY/utils/db"
)

type UserAddressRepository struct {
}

// UserAddres 类型
type UserAddress struct {
	Id        uint   `json:"id"`
	UserId    string `json:"userId"`
	Receiver  string `json:"receiver"`
	Mobile    string `json:"mobile"`
	Province  string `json:"province"`
	City      string `json:"city"`
	District  string `json:"district"`
	Detail    string `json:"detail"`
	Extend    string `json:"extend"`
	IsDefault int    `json:"isDefault"`
}

// NewUserAddressRepository 实例化 DAO
func NewUserAddressRepository() *UserAddressRepository {
	return &UserAddressRepository{}
}

func (r *UserAddressRepository) Create(t *model.UserAddress) (*model.UserAddress, error) {
	err := db.GetMysql().Create(t).Error
	return t, err
}

// DeleteById 删除用户地址
func (r *UserAddressRepository) DeleteById(id uint) error {
	if err := db.GetMysql().Where("id = ?", id).Delete(UserAddress{}).Error; err != nil {
		return err
	}

	return nil
}

// UpdateByUserId 修改用户地址
func (r *UserAddressRepository) UpdateByUserId(id uint, t *model.UserAddress) (*model.UserAddress, error) {
	var ret = new(model.UserAddress)
	err := db.GetMysql().Model(&ret).Where("id=?", id).Updates(t).Error
	return ret, err
}

// GetByUserId 获取用户地址
func (r *UserAddressRepository) GetByUserId(id uint, t *model.UserAddress) (*model.UserAddress, error) {
	var ret = new(model.UserAddress)
	err := db.GetMysql().Model(&ret).Where("id=?", id).Updates(t).Error
	return ret, err
}

// Get 获取用户列表
func (r *UserAddressRepository) Get(page int, pageSize int, username string) ([]model.UserAddress, error) {
	var usersAdress []model.UserAddress
	var err error
	if pageSize < 1 {
		pageSize = 10
	}
	if page < 1 {
		page = 1
	}
	if username != "" {
		err = db.GetMysql().Where("username like ?", "%"+username+"%").Limit(pageSize).Offset((page - 1) * pageSize).Find(&usersAdress).Error
	} else {
		err = db.GetMysql().Limit(pageSize).Offset((page - 1) * pageSize).Find(&usersAdress).Error
	}
	return usersAdress, err
}
