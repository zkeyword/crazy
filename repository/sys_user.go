package repository

import (
	"CRAZY/model"
	"CRAZY/utils/db"
	"time"
)

type UserRepository struct {
}

// User response
type UserList struct {
	Data     []model.User
	Total    int
	PageSize int
	Page     int
}

// User 类型
type User struct {
	ID        uint
	Title     string
	Content   string
	AuthorID  uint
	UpdatedAt time.Time
	TagID     int
	TagName   string
}

// NewUserRepository 实例化 DAO
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Get 获取用户
func (r *UserRepository) Get(id int64) *User {
	ret := &User{}

	if err := db.GetMysql().First(ret, "id = ?", id).Error; err != nil {
		return nil
	}

	return ret
}

func (r *UserRepository) FindByID(id int64) (*model.User, error) {
	var user = new(model.User)
	result := db.GetMysql().Where("id=?", id).Find(user)
	err := result.Error
	return user, err
}

func (r *UserRepository) FindByIDs(ids []uint) ([]model.User, error) {
	var users = make([]model.User, 0)
	err := db.GetMysql().Where("id in (?)", ids).Find(&users).Error
	for idx := range users {
		users[idx].Password = ""
	}
	return users, err
}

// Create 创建用户
func (r *UserRepository) Create(t *model.User) (uint, error) {
	err := db.GetMysql().Create(t).Error
	return t.ID, err
}
