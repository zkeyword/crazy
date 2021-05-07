package services

import (
	"CRAZY/model"
	"CRAZY/repository"
	"time"
)

type RoleService interface {
	Get(id uint) *repository.Role
	Create(Role *model.Role, PermissionKeys string) (*ReturnPolePermission, error)
	UpdateById(id uint, Role *model.Role, PermissionKeys string) (*ReturnPolePermission, error)
	DeleteById(id uint) error
}

type roleService struct {
	repo           *repository.RoleRepository
	rolePermission *repository.RolePermissionRepository
}

type ReturnPolePermission struct {
	ID             uint      `json:"id"`
	Name           string    `json:"name"`
	PermissionKeys string    `json:"permissionKeys"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

var NewRoleService = newRoleService()

func newRoleService() RoleService {
	return &roleService{
		repo: repository.NewRoleRepository(),
	}
}

func (s *roleService) Get(id uint) *repository.Role {
	return s.repo.Get(id)
}

func (s *roleService) Create(Role *model.Role, PermissionKeys string) (*ReturnPolePermission, error) {
	ret, err := s.repo.Create(Role)
	if err == nil {
		s.rolePermission.Create(ret.ID, PermissionKeys)
	}
	returnValue := &ReturnPolePermission{
		ID:             ret.ID,
		Name:           ret.Name,
		PermissionKeys: PermissionKeys,
		CreatedAt:      ret.CreatedAt,
		UpdatedAt:      ret.UpdatedAt,
	}
	return returnValue, err
}

func (s *roleService) UpdateById(id uint, Role *model.Role, PermissionKeys string) (*ReturnPolePermission, error) {
	ret, err := s.repo.UpdateById(id, Role)
	if err == nil {
		s.rolePermission.UpdateById(id, PermissionKeys)
	} else {
		PermissionKeys = ""
	}
	returnValue := &ReturnPolePermission{
		ID:             ret.ID,
		Name:           ret.Name,
		PermissionKeys: PermissionKeys,
		CreatedAt:      ret.CreatedAt,
		UpdatedAt:      ret.UpdatedAt,
	}
	return returnValue, err
}

func (s *roleService) DeleteById(id uint) error {
	err := s.repo.DeleteById(id)
	return err
}
