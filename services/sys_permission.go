package services

import (
	"CRAZY/model"
	"CRAZY/repository"
)

type PermissionService interface {
	Get(id int64) *repository.Permission
	Create(Article *model.Permission) (uint, error)
}

type permissionService struct {
	repo *repository.PermissionRepository
}

var NewPermissionService = newPermissionService()

func newPermissionService() PermissionService {
	return &permissionService{
		repo: repository.NewPermissionRepository(),
	}
}

func (s *permissionService) Get(id int64) *repository.Permission {
	return s.repo.Get(id)
}

func (s *permissionService) Create(Permission *model.Permission) (uint, error) {
	ID, err := s.repo.Create(Permission)
	return ID, err
}
