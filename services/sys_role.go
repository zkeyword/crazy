package services

import (
	"CRAZY/model"
	"CRAZY/repository"
	"time"
)

type RoleService interface {
	Get(page int, pageSize int, name string) (*ReturnRoleList, error)
	GetById(id uint) *repository.Role
	Create(Role *model.Role, PermissionKeys string) (*ReturnPolePermission, error)
	UpdateById(id uint, Role *model.Role, PermissionKeys string) (*ReturnPolePermission, error)
	DeleteById(id uint) error
}

type roleService struct {
	repo           *repository.RoleRepository
	rolePermission *repository.RolePermissionRepository
}

type ReturnRoleList struct {
	Page     int          `json:"page"`
	PageSize int          `json:"pageSize"`
	Total    int          `json:"total"`
	List     []model.Role `json:"list"`
}

type ReturnPolePermission struct {
	ID             uint
	Name           string
	PermissionKeys string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

var NewRoleService = newRoleService()

func newRoleService() RoleService {
	return &roleService{
		repo: repository.NewRoleRepository(),
	}
}

func (s *roleService) Get(page int, pageSize int, name string) (*ReturnRoleList, error) {
	ret, err := s.repo.Get(page, pageSize, name)
	count := s.repo.GetRoleCount(name)
	returnValue := &ReturnRoleList{
		Page:     page,
		PageSize: pageSize,
		Total:    count,
		List:     ret,
	}
	return returnValue, err
}

func (s *roleService) GetById(id uint) *repository.Role {
	return s.repo.GetById(id)
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
