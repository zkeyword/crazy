package sysPermissionService

import (
	"CRAZY/model/system"
	repository "CRAZY/repository/system"
)

var repo = getRepo()

func getRepo() *repository.PermissionRepository {
	return repository.NewPermissionRepository()
}

func Create(Permission *system.Permission) (*system.Permission, error) {
	ret, err := repo.Create(Permission)
	return ret, err
}

func UpdateById(id uint, Permission *system.Permission) (*system.Permission, error) {
	ret, err := repo.UpdateById(id, Permission)
	return ret, err
}

func DeleteById(id uint, key string) error {
	err := repo.DeleteById(id, key)
	return err
}

func GetById(id uint) *repository.Permission {
	return repo.GetById(id)
}

func Get() []system.Permission {
	return repo.Get()
}

type TreeList struct {
	ID       uint        `json:"id"`
	Name     string      `json:"name"`
	Key      string      `json:"key"`
	Status   int         `json:"status"`
	PID      uint        `json:"pid"`
	Children []*TreeList `json:"children"`
}

func GetTree(pid uint) []*TreeList {
	ret := repo.Get()
	return recursion(pid, ret, true)
}

func recursion(pid uint, list []system.Permission, isRoot bool) []*TreeList {
	treeList := []*TreeList{}
	for _, v := range list {
		var isPass bool
		if isRoot {
			isPass = v.ID == pid
		} else {
			isPass = v.PID == pid
		}
		if isPass && v.ID != v.PID {
			child := recursion(v.ID, list, false)
			node := &TreeList{
				ID:     v.ID,
				Name:   v.Name,
				Key:    v.Key,
				Status: v.Status,
				PID:    v.PID,
			}
			node.Children = child
			treeList = append(treeList, node)
		}
	}
	return treeList
}
