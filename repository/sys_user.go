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
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
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
func (r *UserRepository) DeleteById(id uint) error {
	if err := db.GetMysql().Where("id = ?", id).Delete(User{}).Error; err != nil {
		return err
	}

	return nil
}

// UpdateById 修改用户
func (r *UserRepository) UpdateById(id uint, t *model.User) (*model.User, error) {
	var ret = new(model.User)
	err := db.GetMysql().Model(&ret).Where("id=?", id).Updates(t).Error
	return ret, err
}

// Get 获取用户列表
func (r *UserRepository) Get(page int, pageSize int, username string) ([]model.User, error) {
	var users []model.User
	var err error
	if pageSize < 1 {
		pageSize = 10
	}
	if username != "" {
		err = db.GetMysql().Where("username like ?", "%"+username+"%").Limit(pageSize).Offset((page - 1) * pageSize).Find(&users).Error
	} else {
		err = db.GetMysql().Limit(pageSize).Offset((page - 1) * pageSize).Find(&users).Error
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
func (r *UserRepository) GetById(id uint) (*model.User, error) {
	var ret = &model.User{}
	err := db.GetMysql().First(ret, "id = ?", id).Error
	return ret, err
}

type ReturnUser struct {
	ID             uint     `json:"id"`
	RoleID         uint     `json:"roleIDs"`
	PermissionKeys string   `json:"permissionKeys"`
	RoleName       []string `json:"roles"`
}

// GetUserRolePermissionByUserId 获取用户角色关联权限
func (r *UserRepository) GetUserRolePermissionByUserId(id uint) *ReturnUser {
	ret := &ReturnUser{}
	var ret2 []ReturnUser

	err := db.GetMysql().Table("user_roles").Where("user_id = ?", id).Find(&ret2).Error

	ret.RoleID = id

	var roleIDs []uint
	for _, v := range ret2 {
		roleIDs = append(roleIDs, v.RoleID)
	}

	// 获取关联角色
	var role []Role
	db.GetMysql().Table("roles").Where("id IN (?)", roleIDs).Find(&role)

	roleName := make([]string, 0)
	for _, v := range role {
		roleName = append(roleName, v.Name)
	}
	ret.RoleName = roleName

	// 获取关联权限
	var permission []RolePermission
	db.GetMysql().Table("role_permissions").Where("id IN (?)", roleIDs).Find(&permission)

	permissionKey := make([]string, 0)
	for _, v := range permission {
		permissionKey = append(permissionKey, v.PermissionKeys)
	}

	permissionKeys := strings.Split(strings.Join(permissionKey, ","), ",") // 获取权限总集
	ret.PermissionKeys = strings.Join(utils.RemoveRepeated(permissionKeys), ",")

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
