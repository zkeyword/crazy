package repository

import (
	"CRAZY/model"
	"CRAZY/utils"
	"CRAZY/utils/db"
	"strings"
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

type ReturnUser struct {
	ID             uint   `json:"id"`
	Username       string `json:"username"`
	Password       string `json:"password"`
	RoleIDs        string `json:"roleIDs"`
	PermissionKeys string `json:"permissionKeys"`
}

// GetById 获取用户
func (r *UserRepository) Get(page int, pageSize int, username string) ([]model.User, error) {
	var users []model.User
	var err error
	if username != "" {
		err = db.GetMysql().Where("username like ?", "%"+username+"%").Limit(pageSize).Offset((page - 1) * 10).Find(&users).Error
	} else {
		err = db.GetMysql().Limit(pageSize).Offset((page - 1) * 10).Find(&users).Error
	}
	return users, err
}

func (r *UserRepository) GetUserCount(username string) int {
	var users []model.User
	var count int
	if username != "" {
		db.GetMysql().Where("username like ?", "%"+username+"%").Find(&users).Select("count(id)").Count(&count)
	} else {
		db.GetMysql().Find(&users).Select("count(id)").Count(&count)
	}
	return count
}

// GetById 获取用户
func (r *UserRepository) GetById(id int64) *ReturnUser {
	ret := &ReturnUser{}

	err := db.GetMysql().
		// Debug().
		Table("users").
		Select("users.id, users.username, users.password, users.level, users.updated_at, user_roles.role_ids").
		Joins("left join user_roles on users.id = user_roles.user_id").
		// Joins("left join roles on roles.id = user_roles.role_ids").
		Where("users.id = ?", id).
		Find(&ret).
		Error

	roleIDs := strings.Split(ret.RoleIDs, ",")

	// var role []Role
	// db.GetMysql().Table("roles").Where("id IN (?)", roleIDs).Find(&role)

	var permission []RolePermission
	db.GetMysql().Table("role_permissions").Where("id IN (?)", roleIDs).Find(&permission)

	result := make([]string, 0)
	for _, v := range permission {
		result = append(result, v.PermissionKeys)
	}

	tmp := strings.Split(strings.Join(result, ","), ",") // 获取总集
	ret.PermissionKeys = strings.Join(utils.RemoveRepeated(tmp), ",")

	if err != nil {
		return nil
	}

	return ret
}

// GetByName
func (r *UserRepository) GetByUserName(username string) (*model.User, error) {
	var ret = &model.User{}
	err := db.GetMysql().First(ret, "username = ?", username).Error
	return ret, err
}
