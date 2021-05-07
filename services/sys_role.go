package services

import (
	"CRAZY/model"
	"CRAZY/repository"
)

type RoleService interface {
	Get(id int64) *repository.Role
	Create(Role *model.Role, PermissionKeys string) (*model.Role, error)
	UpdateById(id int64, Role *model.Role, PermissionKeys string) (*model.Role, error)
	DeleteById(id int64) error
}

type roleService struct {
	repo           *repository.RoleRepository
	rolePermission *repository.RolePermissionRepository
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

func (s *roleService) Create(Role *model.Role, PermissionKeys string) (*model.Role, error) {
	ret, err := s.repo.Create(Role)
	if err == nil {
		s.rolePermission.Create(ret.ID, PermissionKeys)
	}
	return ret, err
}

func (s *roleService) UpdateById(id int64, Role *model.Role, PermissionKeys string) (*model.Role, error) {
	ret, err := s.repo.UpdateById(id, Role)
	if err == nil {
		s.rolePermission.UpdateById(id, PermissionKeys)
	}
	return ret, err
}

func (s *roleService) DeleteById(id int64) error {
	err := s.repo.DeleteById(id)
	return err
}
