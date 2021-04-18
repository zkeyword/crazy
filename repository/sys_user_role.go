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

func (r *UserRoleRepository) Create(UserId uint, RoleId uint) (*model.UserRole, error) {
	var ret = new(model.UserRole)
	ret.RoleID = RoleId
	ret.UserID = UserId
	err := db.GetMysql().Create(ret).Error
	return ret, err
}

func (r *UserRoleRepository) DeleteById(UserID int64) error {
	if err := db.GetMysql().Where("user_id = ?", UserID).Delete(UserRole{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRoleRepository) UpdateById(UserID int64, RoleID uint) (*model.UserRole, error) {
	var ret = new(model.UserRole)
	data := &UserRole{}
	data.RoleID = RoleID
	err := db.GetMysql().Model(&ret).Where("user_id=?", UserID).Updates(data).Error
	return ret, err
}

func (r *UserRoleRepository) Get(id int64) *UserRole {
	ret := &UserRole{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}
