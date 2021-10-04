package sysUserService

import (
	"CRAZY/model"
	"CRAZY/repository"
)

var userRepo = getUserRepo()

var userRoleRepo = getUserRoleRepo()

func getUserRepo() *repository.UserRepository {
	return repository.NewUserRepository()
}

func getUserRoleRepo() *repository.UserRoleRepository {
	return repository.NewUserRoleRepository()
}

type ReturnUserList struct {
	Page     int          `json:"page"`
	PageSize int          `json:"pageSize"`
	Total    int          `json:"total"`
	List     []model.User `json:"list"`
}

func Get(page int, pageSize int, username string) (*ReturnUserList, error) {
	ret, err := userRepo.Get(page, pageSize, username)
	count := userRepo.GetUserCount(username)
	returnValue := &ReturnUserList{
		Page:     page,
		PageSize: pageSize,
		Total:    count,
		List:     ret,
	}
	return returnValue, err
}

func GetById(id int64) *repository.ReturnUser {
	return userRepo.GetById(id)
}

func Create(User *model.User, roleIds string) (*model.User, error) {
	ret, err := userRepo.Create(User)
	if err == nil {
		userRoleRepo.Create(ret.ID, roleIds)
	}
	return ret, err
}

func PutUserById(id int64, User *model.User, roleIds string) (*model.User, error) {
	ret, err := userRepo.UpdateById(id, User)
	if err == nil {
		userRoleRepo.UpdateById(id, roleIds)
	}
	return ret, err
}

func DeleteById(id int64) error {
	err := userRepo.DeleteById(id)
	if err == nil {
		userRoleRepo.DeleteById(id)
	}
	return err
}

func GetByUserName(username string) (*model.User, error) {
	return userRepo.GetByUserName(username)
}
