package sysRoleService

import (
	"CRAZY/model"
	"CRAZY/repository"
	"time"
)

var roleRepo = getRoleRepo()

var rolePermission = getRolePermission()

func getRoleRepo() *repository.RoleRepository {
	return repository.NewRoleRepository()
}

func getRolePermission() *repository.RolePermissionRepository {
	return repository.NewRolePermissionRepository()
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

func Get(page int, pageSize int, name string) (*ReturnRoleList, error) {
	ret, err := roleRepo.Get(page, pageSize, name)
	count := roleRepo.GetRoleCount(name)
	returnValue := &ReturnRoleList{
		Page:     page,
		PageSize: pageSize,
		Total:    count,
		List:     ret,
	}
	return returnValue, err
}

func GetById(id uint) *repository.Role {
	return roleRepo.GetById(id)
}

func Create(Role *model.Role, PermissionKeys string) (*ReturnPolePermission, error) {
	ret, err := roleRepo.Create(Role)
	if err == nil {
		rolePermission.Create(ret.ID, PermissionKeys)
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

func UpdateById(id uint, Role *model.Role, PermissionKeys string) (*ReturnPolePermission, error) {
	ret, err := roleRepo.UpdateById(id, Role)
	if err == nil {
		rolePermission.UpdateById(id, PermissionKeys)
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

func DeleteById(id uint) error {
	err := roleRepo.DeleteById(id)
	return err
}