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

func (r *UserRoleRepository) Create(userId uint, username string, roleId uint) (*model.UserRole, error) {
	var ret = new(model.UserRole)
	ret.RoleID = roleId
	ret.UserID = userId
	ret.Username = username
	_db, err := db.GetMysql()
	if err != nil {
		return ret, err
	}
	err = _db.Create(ret).Error
	return ret, err
}

func (r *UserRoleRepository) DeleteByUserId(userID uint) error {
	_db, err := db.GetMysql()
	if err != nil {
		return err
	}
	err = _db.Where("user_id = ?", userID).Delete(UserRole{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRoleRepository) DeleteByRoleId(roleID uint) error {
	_db, err := db.GetMysql()
	if err != nil {
		return err
	}
	err = _db.Where("role_id = ?", roleID).Delete(UserRole{}).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRoleRepository) DeleteByRoleIdAndUserId(userID uint, roleID uint) error {
	_db, err := db.GetMysql()
	if err != nil {
		return nil
	}
	err = _db.Where("user_id = ? AND role_id = ?", userID, roleID).Delete(UserRole{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRoleRepository) GetByRoleID(id uint) *[]UserRole {
	ret := &[]UserRole{}
	_db, err := db.GetMysql()
	if err != nil {
		return nil
	}
	err = _db.Find(ret, "role_id = ?", id).Error
	if err != nil {
		return nil
	}
	return ret
}

func (r *UserRoleRepository) GetByUserID(id uint) *[]UserRole {
	ret := &[]UserRole{}
	_db, err := db.GetMysql()
	if err != nil {
		return nil
	}
	err = _db.Find(ret, "user_id = ?", id).Error
	if err != nil {
		return nil
	}
	return ret
}
