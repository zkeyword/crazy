package repository

import (
	"CRAZY/model"
	"CRAZY/utils/db"
)

type UserRoleRepository struct {
}

// User 类型
type UserRole struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"userID"`
	RoleID   uint   `json:"roleID"`
	Username string `json:"username"`
}

// NewUserRoleRepository 实例化 DAO
func NewUserRoleRepository() *UserRoleRepository {
	return &UserRoleRepository{}
}

func (r *UserRoleRepository) Create(userId uint, roleId uint, username string) (*model.UserRole, error) {
	var ret = new(model.UserRole)
	ret.RoleID = roleId
	ret.UserID = userId
	ret.Username = username
	err := db.GetMysql().Create(ret).Error
	return ret, err
}

func (r *UserRoleRepository) DeleteById(userID int64) error {
	if err := db.GetMysql().Where("user_id = ?", userID).Delete(UserRole{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRoleRepository) GetByRoleID(id uint) *[]UserRole {
	ret := &[]UserRole{}

	if err := db.GetMysql().Find(ret, "role_id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}

func (r *UserRoleRepository) GetByUserID(id uint) *[]UserRole {
	ret := &[]UserRole{}

	if err := db.GetMysql().First(ret, "role_id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}
