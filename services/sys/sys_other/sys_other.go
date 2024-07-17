package otherService

import (
	"CRAZY/model/system"
	repository "CRAZY/repository/system"
)

var repo = getRepo()

func getRepo() *repository.OtherRepository {
	return repository.NewOtherRepository()
}

func Get() ([]system.Other, error) {
	ret, err := repo.Get()
	return ret, err
}

func Create(Other *system.Other) (*system.Other, error) {
	ret, err := repo.Create(Other)
	return ret, err
}

func BatchCreate(Other []*system.Other) ([]*system.Other, error) {
	ret, err := repo.BatchCreate(Other)
	return ret, err
}

func PutById(id uint, Other *system.Other) (*system.Other, error) {
	ret, err := repo.UpdateById(id, Other)
	return ret, err
}

func DeleteById(id uint) error {
	err := repo.DeleteById(id)
	return err
}

func GetById(id uint) (*system.Other, error) {
	ret, err := repo.GetById(id)
	return ret, err
}
