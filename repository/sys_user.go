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
// func (r *UserRepository) Create(t *model.User, RoleId uint) (*model.User, error) {
// 	tx := db.GetMysql().Begin()
// 	c := tx.Create(t)
// 	rowsAffected := c.RowsAffected
// 	if rowsAffected == 0 {
// 		tx.Rollback()
// 		return t, c.Error
// 	}
// 	var ur *UserRoleRepository
// 	_, err := ur.Create(t.ID, RoleId)
// 	if err != nil {
// 		tx.Rollback()
// 	} else {
// 		tx.Commit()
// 	}
// 	return t, err
// }
func (r *UserRepository) Create(t *model.User) (*model.User, error) {
	err := db.GetMysql().Create(t).Error
	return t, err
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

	err := db.GetMysql().
		// Debug().
		Table("users a").
		Select("a.id, a.username, a.password, a.level, a.updated_at, t.id role_id, t.name role_name").
		Joins("left join user_roles r on a.id = r.user_id").
		Joins("left join roles t on t.id = r.role_id").
		Where("a.id = ?", id).
		Find(&ret).
		Error

	if err != nil {
		return nil
	}

	return ret
}
