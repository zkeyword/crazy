package sysUserService

import (
	"CRAZY/model"
	"CRAZY/repository"
	"CRAZY/utils"
	"strings"
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

func GetById(id uint) (*model.User, error) {
	return userRepo.GetById(id)
}

func GetUserRolePermissionByUserId(id uint) *repository.ReturnUser {
	return userRepo.GetUserRolePermissionByUserId(id)
}

func Create(User *model.User, roleIds string) (*model.User, error) {
	ret, err := userRepo.Create(User)
	roleIdArr := strings.Split(roleIds, ",")
	for _, v := range roleIdArr {
		userRoleRepo.Create(ret.ID, ret.Username, utils.StrToUInt(v))
	}
	return ret, err
}

func PutUserById(id uint, User *model.User, roleIds string) (*model.User, error) {
	ret, err := userRepo.UpdateById(id, User)
	roleIdArr := strings.Split(roleIds, ",")
	for _, v := range roleIdArr {
		userRoleRepo.DeleteByUserId(id)
		userRoleRepo.Create(ret.ID, ret.Username, utils.StrToUInt(v))
	}
	return ret, err
}

func DeleteById(id uint) error {
	err := userRepo.DeleteById(id)
	if err == nil {
		userRoleRepo.DeleteByUserId(id)
	}
	return err
}

func GetByUserName(username string) (*model.User, error) {
	return userRepo.GetByUserName(username)
}
