package services

import (
	"CRAZY/model"
	"CRAZY/repository"
)

type RoleService interface {
	Get(id int64) *repository.Role
	Create(Role *model.Role) (uint, error)
	UpdateById(id int64, Role *model.Role) (*model.Role, error)
	DeleteById(id int64) error
}

type roleService struct {
	repo *repository.RoleRepository
}

var NewRoleService = newRoleService()

func newRoleService() RoleService {
	return &roleService{
		repo: repository.NewRoleRepository(),
	}
}

func (s *roleService) Get(id int64) *repository.Role {
	return s.repo.Get(id)
}

func (s *roleService) Create(Role *model.Role) (uint, error) {
	ID, err := s.repo.Create(Role)
	return ID, err
}

func (s *roleService) UpdateById(id int64, Role *model.Role) (*model.Role, error) {
	ret, err := s.repo.UpdateById(id, Role)
	return ret, err
}

func (s *roleService) DeleteById(id int64) error {
	err := s.repo.DeleteById(id)
	return err
}
