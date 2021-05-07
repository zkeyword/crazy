package services

import (
	"CRAZY/model"
	"CRAZY/repository"
	"time"
)

type PermissionService interface {
	Get(id int64) *ReturnPermission
	Create(Article *model.Permission) (*model.Permission, error)
	UpdateById(id int64, Permission *model.Permission) (*model.Permission, error)
	DeleteById(id int64) error
}

type permissionService struct {
	repo *repository.PermissionRepository
}

type ReturnPermission struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Key       string    `json:"key"`
	Status    int       `json:"status"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

var NewPermissionService = newPermissionService()

func newPermissionService() PermissionService {
	return &permissionService{
		repo: repository.NewPermissionRepository(),
	}
}

func (s *permissionService) Get(id int64) *ReturnPermission {
	ret := s.repo.Get(id)
	returnValue := &ReturnPermission{
		ID:        ret.ID,
		Name:      ret.Name,
		Key:       ret.Key,
		Status:    ret.Status,
		CreatedAt: ret.CreatedAt,
		UpdatedAt: ret.UpdatedAt,
	}
	return returnValue
}

func (s *permissionService) Create(Permission *model.Permission) (*model.Permission, error) {
	ret, err := s.repo.Create(Permission)
	return ret, err
}

func (s *permissionService) UpdateById(id int64, Permission *model.Permission) (*model.Permission, error) {
	ret, err := s.repo.UpdateById(id, Permission)
	return ret, err
}

func (s *permissionService) DeleteById(id int64) error {
	err := s.repo.DeleteById(id)
	return err
}
