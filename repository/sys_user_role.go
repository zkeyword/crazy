package repository

import (
	"CRAZY/utils/db"
)

type UserRoleRepository struct {
}

// User 类型
type UserRole struct {
	ID     uint
	UserId string
	RoleId string
}

// NewUserRoleRepository 实例化 DAO
func NewUserRoleRepository() *UserRoleRepository {
	return &UserRoleRepository{}
}

func (r *UserRoleRepository) Create(UserId string, RoleId string) (*UserRole, error) {
	ret := &UserRole{}
	ret.RoleId = RoleId
	ret.UserId = UserId
	err := db.GetMysql().Create(ret).Error
	return ret, err
}

func (r *UserRoleRepository) DeleteById(id int64) error {
	if err := db.GetMysql().Where("id = ?", id).Delete(UserRole{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRoleRepository) Get(id int64) *UserRole {
	ret := &UserRole{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}
