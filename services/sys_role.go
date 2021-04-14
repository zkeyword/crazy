package services

import (
	"CRAZY/model"
	"CRAZY/repository"
)

type RoleService interface {
	Get(id int64) *repository.Role
	Create(Article *model.Role) (uint, error)
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
