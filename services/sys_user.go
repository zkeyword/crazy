package services

import (
	"CRAZY/model"
	"CRAZY/repository"
)

// UserService user服务
type UserService interface {
	Get(id int) *repository.User
	FindByID(userID int) (*model.User, error)
	// FindByIDs(ids []uint) []model.User
	Create(Article *model.User) (uint, error)
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

func (s *userService) Get(id int) *repository.User {
	return s.repo.Get(id)
}

func (s *userService) FindByID(userID int) (*model.User, error) {
	user, err := s.repo.FindByID(userID)
	return user, err
}

// func (s *userService) FindByIDs(ids []uint) []model.User {
// 	users, _ := s.repo.FindByIDs(ids)
// 	return users
// }

func (s *userService) Create(User *model.User) (uint, error) {
	ID, err := s.repo.Create(User)
	return ID, err
}
