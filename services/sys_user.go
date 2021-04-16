package services

import (
	"CRAZY/model"
	"CRAZY/repository"
	"strconv"
)

// UserService user服务
type UserService interface {
	Get(id int64) *repository.User
	Create(User *model.User, RoleId uint) (*model.User, error)
	DeleteById(id int64) error
	PutUserById(id int64, User *model.User) (*model.User, error)
}

type userService struct {
	repo  *repository.UserRepository
	repo2 *repository.UserRoleRepository
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

func (s *userService) Create(User *model.User, RoleId uint) (*model.User, error) {
	ret, err := s.repo.Create(User)
	s.repo2.Create(strconv.Itoa(int(ret.ID)), strconv.Itoa(int(RoleId)))
	return ret, err
}

func (s *userService) PutUserById(id int64, User *model.User) (*model.User, error) {
	ret, err := s.repo.UpdateById(id, User)
	return ret, err
}

func (s *userService) DeleteById(id int64) error {
	err := s.repo.DeleteById(id)
	return err
}
