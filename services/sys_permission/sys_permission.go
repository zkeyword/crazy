package sysPermissionService

import (
	"CRAZY/model"
	"CRAZY/repository"
)

var repo = getRepo()

func getRepo() *repository.PermissionRepository {
	return repository.NewPermissionRepository()
}

func Get(id int64) *repository.Permission {
	return repo.Get(id)
}

func Create(Permission *model.Permission) (*model.Permission, error) {
	ret, err := repo.Create(Permission)
	return ret, err
}

func UpdateById(id int64, Permission *model.Permission) (*model.Permission, error) {
	ret, err := repo.UpdateById(id, Permission)
	return ret, err
}

func DeleteById(id int64) error {
	err := repo.DeleteById(id)
	return err
}
