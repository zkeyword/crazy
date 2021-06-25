package services

import (
	"CRAZY/model"
	"CRAZY/repository"
)

// UserService user服务
type UserService interface {
	Get(page int, pageSize int, username string) (*ReturnUserList, error)
	GetById(id int64) *repository.ReturnUser
	GetByUserName(username string) (*model.User, error)
	Create(User *model.User, roleIds string) (*model.User, error)
	DeleteById(id int64) error
	PutUserById(id int64, User *model.User, roleIds string) (*model.User, error)
}

type userRepository struct {
	repo     *repository.UserRepository
	userRole *repository.UserRoleRepository
}

type ReturnUserList struct {
	Page     int          `json:"page"`
	PageSize int          `json:"pageSize"`
	Total    int          `json:"total"`
	List     []model.User `json:"list"`
}

// NewArticleService 实例化ArticleService
var NewUserService = newUserService()

func newUserService() UserService {
	return &userRepository{
		repo: repository.NewUserRepository(),
	}
}

func (s *userRepository) Get(page int, pageSize int, username string) (*ReturnUserList, error) {
	ret, err := s.repo.Get(page, pageSize, username)
	count := s.repo.GetUserCount(username)
	returnValue := &ReturnUserList{
		Page:     page,
		PageSize: pageSize,
		Total:    count,
		List:     ret,
	}
	return returnValue, err
}

func (s *userRepository) GetById(id int64) *repository.ReturnUser {
	return s.repo.GetById(id)
}

// func (s *userRepository) Create(User *model.User, RoleId uint) (*model.User, error) {
// 	ret, err := s.repo.Create(User, RoleId)
// 	// s.userRole.Create(ret.ID, RoleId)
// 	return ret, err
// }
func (s *userRepository) Create(User *model.User, roleIds string) (*model.User, error) {
	ret, err := s.repo.Create(User)
	if err == nil {
		s.userRole.Create(ret.ID, roleIds)
	}
	return ret, err
}

func (s *userRepository) PutUserById(id int64, User *model.User, roleIds string) (*model.User, error) {
	ret, err := s.repo.UpdateById(id, User)
	if err == nil {
		s.userRole.UpdateById(id, roleIds)
	}
	return ret, err
}

func (s *userRepository) DeleteById(id int64) error {
	err := s.repo.DeleteById(id)
	if err == nil {
		s.userRole.DeleteById(id)
	}
	return err
}

func (s *userRepository) GetByUserName(username string) (*model.User, error) {
	return s.repo.GetByUserName(username)
}
