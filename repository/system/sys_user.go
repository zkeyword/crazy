package system

import (
	"CRAZY/model/system"
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
//
//	func (r *UserRepository) Create(t *system.SystemUser, RoleId uint) (*system.SystemUser, error) {
//		tx := db.GetMysql().Begin()
//		c := tx.Create(t)
//		rowsAffected := c.RowsAffected
//		if rowsAffected == 0 {
//			tx.Rollback()
//			return t, c.Error
//		}
//		var ur *UserRoleRepository
//		_, err := ur.Create(t.ID, RoleId)
//		if err != nil {
//			tx.Rollback()
//		} else {
//			tx.Commit()
//		}
//		return t, err
//	}
func (r *UserRepository) Create(t *system.SystemUser) (*system.SystemUser, error) {
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Create(t).Error
	return t, err
}

// DeleteById 删除用户
func (r *UserRepository) DeleteById(id uint) error {
	_db, err := db.GetMysql()
	if err != nil {
		return err
	}
	return _db.Where("id = ?", id).Delete(User{}).Error
}

// UpdateById 修改用户
func (r *UserRepository) UpdateById(id uint, t *system.SystemUser) (*system.SystemUser, error) {
	var ret = new(system.SystemUser)
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Model(&ret).Where("id=?", id).Updates(t).Error
	if err == nil {
		ret.ID = id
	}
	return ret, err
}

// Get 获取用户列表
func (r *UserRepository) Get(page int, pageSize int, username string) ([]system.SystemUser, error) {
	var users []system.SystemUser
	var err error
	if pageSize < 1 {
		pageSize = 10
	}
	if page < 1 {
		page = 1
	}
	_db, err := db.GetMysql()
	if err != nil {
		return users, err
	}
	if username != "" {
		err = _db.Where("username like ?", "%"+username+"%").Limit(pageSize).Offset((page - 1) * pageSize).Find(&users).Error
	} else {
		err = _db.Limit(pageSize).Offset((page - 1) * pageSize).Find(&users).Error
	}
	return users, err
}

func (r *UserRepository) GetUserCount(username string) int64 {
	var users []system.SystemUser
	var count int64
	_db, err := db.GetMysql()
	if err != nil {
		return 0
	}
	if username != "" {
		_db.Where("username like ?", "%"+username+"%").Find(&users).Select("count(id)").Count(&count)
	} else {
		_db.Find(&users).Select("count(id)").Count(&count)
	}
	return count
}

// GetById 获取用户
func (r *UserRepository) GetById(id uint) (*system.SystemUser, error) {
	var ret = &system.SystemUser{}
	_db, err := db.GetMysql()
	if err != nil {
		return ret, err
	}
	err = _db.First(ret, "id = ?", id).Error
	return ret, err
}

// GetLoginStatusById 获取用户登录状态 // 使用redis后冗余
func (r *UserRepository) GetLoginStatusById(id uint) (int, error) {
	var user system.SystemUser
	_db, err := db.GetMysql()
	if err != nil {
		return 0, err
	}
	err = _db.Select("login_status").First(&user, "id = ?", id).Error
	return user.LoginStatus, err
}

type ReturnRolePermission struct {
	ID             uint             `json:"id"`
	RoleID         uint             `json:"roleIDs"`
	PermissionKeys string           `json:"permissionKeys"`
	RoleName       []string         `json:"roles"`
	Permission     []RolePermission `json:"permissions"`
}

type ReturnUserRole struct {
	ID     uint `json:"id"`
	RoleID uint `json:"roleIDs"`
	UserID uint `json:"userID"`
}

// GetUserRolePermissionByUserId 获取用户角色关联权限
func (r *UserRepository) GetUserRolePermissionByUserId(id uint) *ReturnRolePermission {
	_db, err := db.GetMysql()
	if err != nil {
		return nil
	}

	ret := &ReturnRolePermission{}
	var ret2 []ReturnUserRole

	err = _db.Table("user_roles").Where("user_id = ?", id).Find(&ret2).Error

	ret.RoleID = id

	var roleIDs []uint
	for _, v := range ret2 {
		roleIDs = append(roleIDs, v.RoleID)
	}

	// 获取关联角色
	var role []Role
	_db.Table("roles").Where("id IN (?)", roleIDs).Find(&role)

	roleName := make([]string, 0)
	for _, v := range role {
		roleName = append(roleName, v.Name)
	}
	ret.RoleName = roleName

	// 获取关联权限
	var permission []RolePermission
	_db.Table("role_permissions").Where("role_id IN (?)", roleIDs).Find(&permission)

	ret.Permission = permission

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
func (r *UserRepository) GetByUserName(username string) (*system.SystemUser, error) {
	var ret = &system.SystemUser{}
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.First(ret, "username = ?", username).Error
	return ret, err
}
