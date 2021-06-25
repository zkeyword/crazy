package services

import (
	"CRAZY/model"
	"CRAZY/repository"
)

type PermissionService interface {
	Get(id int64) *repository.Permission
	Create(Article *model.Permission) (*model.Permission, error)
	UpdateById(id int64, Permission *model.Permission) (*model.Permission, error)
	DeleteById(id int64) error
}

type permissionRepository struct {
	repo *repository.PermissionRepository
}

var NewPermissionService = newPermissionService()

func newPermissionService() PermissionService {
	return &permissionRepository{
		repo: repository.NewPermissionRepository(),
	}
}

func (s *permissionRepository) Get(id int64) *repository.Permission {
	return s.repo.Get(id)
}

func (s *permissionRepository) Create(Permission *model.Permission) (*model.Permission, error) {
	ret, err := s.repo.Create(Permission)
	return ret, err
}

func (s *permissionRepository) UpdateById(id int64, Permission *model.Permission) (*model.Permission, error) {
	ret, err := s.repo.UpdateById(id, Permission)
	return ret, err
}

func (s *permissionRepository) DeleteById(id int64) error {
	err := s.repo.DeleteById(id)
	return err
}
