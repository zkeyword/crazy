package repository

import (
	"CRAZY/model"
	"CRAZY/utils/db"
	"time"
)

type RoleRepository struct {
}

type RoleList struct {
	Data     []model.Role
	Total    int
	PageSize int
	Page     int
}

type Role struct {
	ID        uint
	Title     string
	Content   string
	AuthorID  uint
	UpdatedAt time.Time
	TagID     int
	TagName   string
}

func NewRoleRepository() *RoleRepository {
	return &RoleRepository{}
}

// Get 获取用户
func (r *RoleRepository) Get(id int64) *Role {
	ret := &Role{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}

// Create 创建用户
func (r *RoleRepository) Create(t *model.Role) (uint, error) {
	err := db.GetMysql().Create(t).Error
	return t.ID, err
}
