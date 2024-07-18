package shop

import (
	"CRAZY/model/shop"
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

func (r *UserAddressRepository) Create(t *shop.ShopUserAddress) (*shop.ShopUserAddress, error) {
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Create(t).Error
	return t, err
}

// DeleteById 删除用户地址
func (r *UserAddressRepository) DeleteById(id uint) error {
	_db, err := db.GetMysql()
	if err != nil {
		return err
	}
	return _db.Where("id = ?", id).Delete(UserAddress{}).Error
}

// UpdateByUserId 修改用户地址
func (r *UserAddressRepository) UpdateByUserId(id uint, t *shop.ShopUserAddress) (*shop.ShopUserAddress, error) {
	var ret = new(shop.ShopUserAddress)
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Model(&ret).Where("id=?", id).Updates(t).Error
	return ret, err
}

// GetByUserId 获取用户地址
func (r *UserAddressRepository) GetByUserId(id uint, t *shop.ShopUserAddress) (*shop.ShopUserAddress, error) {
	var ret = new(shop.ShopUserAddress)
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Model(&ret).Where("id=?", id).Updates(t).Error
	return ret, err
}

// Get 获取用户列表
func (r *UserAddressRepository) Get(page int, pageSize int, username string) ([]shop.ShopUserAddress, error) {
	var usersAdress []shop.ShopUserAddress
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
	if username != "" {
		err = _db.Where("username like ?", "%"+username+"%").Limit(pageSize).Offset((page - 1) * pageSize).Find(&usersAdress).Error
	} else {
		err = _db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&usersAdress).Error
	}
	return usersAdress, err
}
