package sys

import (
	"CRAZY/model/sys"
	"CRAZY/utils/db"
)

type UserRoleRepository struct {
}

// User 类型
type SysUserRole struct {
	ID       uint   `json:"id"`
	UserID   uint   `json:"userID"`
	RoleID   uint   `json:"roleID"`
	Username string `json:"username"`
}

// NewUserRoleRepository 实例化 DAO
func NewUserRoleRepository() *UserRoleRepository {
	return &UserRoleRepository{}
}

func (r *UserRoleRepository) Create(userId uint, username string, roleId uint) (*sys.SysUserRole, error) {
	var ret = new(sys.SysUserRole)
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
	return _db.Where("user_id = ?", userID).Delete(SysUserRole{}).Error
}

func (r *UserRoleRepository) DeleteByRoleId(roleID uint) error {
	_db, err := db.GetMysql()
	if err != nil {
		return err
	}
	return _db.Where("role_id = ?", roleID).Delete(SysUserRole{}).Error
}

func (r *UserRoleRepository) DeleteByRoleIdAndUserId(userID uint, roleID uint) error {
	_db, err := db.GetMysql()
	if err != nil {
		return nil
	}
	return _db.Where("user_id = ? AND role_id = ?", userID, roleID).Delete(SysUserRole{}).Error
}

func (r *UserRoleRepository) GetByRoleID(id uint) *[]SysUserRole {
	ret := &[]SysUserRole{}
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

func (r *UserRoleRepository) GetByUserID(id uint) *[]SysUserRole {
	ret := &[]SysUserRole{}
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
