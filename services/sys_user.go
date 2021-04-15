package services

import (
	"CRAZY/model"
	"CRAZY/repository"
)

// UserService user服务
type UserService interface {
	Get(id int64) *repository.User
	FindByID(userID int64) (*model.User, error)
	FindByIDs(ids []uint) []model.User
	Create(User *model.User) (uint, error)
	DeleteById(id int64) (*repository.User, error)
	PutUserById(id int64, User *model.User) (*model.User, error)
}

type userService struct {
	repo *repository.UserRepository
}

// NewArticleService 实例化ArticleService
var NewUserService = newUserService()

func newUserService() UserService {
	return &userService{
		repo: repository.NewUserRepository(),
	}
}

func (s *userService) Get(id int64) *repository.User {
	return s.repo.Get(id)
}

func (s *userService) FindByID(userID int64) (*model.User, error) {
	user, err := s.repo.FindByID(userID)
	return user, err
}

func (s *userService) FindByIDs(ids []uint) []model.User {
	users, _ := s.repo.FindByIDs(ids)
	return users
}

func (s *userService) Create(User *model.User) (uint, error) {
	ID, err := s.repo.Create(User)
	return ID, err
}

func (s *userService) PutUserById(id int64, User *model.User) (*model.User, error) {
	ID, err := s.repo.UpdateById(id, User)
	return ID, err
}

func (s *userService) DeleteById(id int64) (*repository.User, error) {
	ID, err := s.repo.DeleteById(id)
	return ID, err
}
