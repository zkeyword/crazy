package repository

import (
	"CRAZY/model"
	"CRAZY/utils/db"
)

type UserRoleRepository struct {
}

// User 类型
type UserRole struct {
	ID     uint
	UserID uint
	RoleID uint
}

// NewUserRoleRepository 实例化 DAO
func NewUserRoleRepository() *UserRoleRepository {
	return &UserRoleRepository{}
}

func (r *UserRoleRepository) Create(userId uint, roleId uint) (*model.UserRole, error) {
	var ret = new(model.UserRole)
	ret.RoleID = roleId
	ret.UserID = userId
	err := db.GetMysql().Create(ret).Error
	return ret, err
}

func (r *UserRoleRepository) DeleteById(userID int64) error {
	if err := db.GetMysql().Where("user_id = ?", userID).Delete(UserRole{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRoleRepository) UpdateById(userID int64, roleID uint) (*model.UserRole, error) {
	var ret = new(model.UserRole)
	data := &UserRole{}
	data.RoleID = roleID
	err := db.GetMysql().Model(&ret).Where("user_id=?", userID).Updates(data).Error
	return ret, err
}

func (r *UserRoleRepository) Get(id int64) *UserRole {
	ret := &UserRole{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}
