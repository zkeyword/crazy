package repository

import (
	"CRAZY/model"
	"CRAZY/utils/db"
	"time"
)

type PermissionRepository struct {
}

type PermissionList struct {
	Data     []model.Permission
	Total    int
	PageSize int
	Page     int
}

type Permission struct {
	ID        uint
	Title     string
	Content   string
	AuthorID  uint
	UpdatedAt time.Time
	TagID     int
	TagName   string
}

func NewPermissionRepository() *PermissionRepository {
	return &PermissionRepository{}
}

// Get 获取用户
func (r *PermissionRepository) Get(id int64) *Permission {
	ret := &Permission{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}

// Create 创建用户
func (r *PermissionRepository) Create(t *model.Permission) (uint, error) {
	err := db.GetMysql().Create(t).Error
	return t.ID, err
}
