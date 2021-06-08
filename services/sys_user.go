package services

import (
	"CRAZY/model"
	"CRAZY/repository"
)

// UserService user服务
type UserService interface {
	Get(page int64) (*model.User, error)
	GetById(id int64) *repository.ReturnUser
	GetByUserName(username string) (*model.User, error)
	Create(User *model.User, roleIds string) (*model.User, error)
	DeleteById(id int64) error
	PutUserById(id int64, User *model.User, roleIds string) (*model.User, error)
}

type userService struct {
	repo     *repository.UserRepository
	userRole *repository.UserRoleRepository
}

// NewArticleService 实例化ArticleService
var NewUserService = newUserService()

func newUserService() UserService {
	return &userService{
		repo: repository.NewUserRepository(),
	}
}

func (s *userService) Get(page int64) (*model.User, error) {
	return s.repo.Get(page)
}

func (s *userService) GetById(id int64) *repository.ReturnUser {
	return s.repo.GetById(id)
}

// func (s *userService) Create(User *model.User, RoleId uint) (*model.User, error) {
// 	ret, err := s.repo.Create(User, RoleId)
// 	// s.userRole.Create(ret.ID, RoleId)
// 	return ret, err
// }
func (s *userService) Create(User *model.User, roleIds string) (*model.User, error) {
	ret, err := s.repo.Create(User)
	if err == nil {
		s.userRole.Create(ret.ID, roleIds)
	}
	return ret, err
}

func (s *userService) PutUserById(id int64, User *model.User, roleIds string) (*model.User, error) {
	ret, err := s.repo.UpdateById(id, User)
	if err == nil {
		s.userRole.UpdateById(id, roleIds)
	}
	return ret, err
}

func (s *userService) DeleteById(id int64) error {
	err := s.repo.DeleteById(id)
	if err == nil {
		s.userRole.DeleteById(id)
	}
	return err
}

func (s *userService) GetByUserName(username string) (*model.User, error) {
	return s.repo.GetByUserName(username)
}
