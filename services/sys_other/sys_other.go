package otherService

import (
	"CRAZY/model"
	"CRAZY/repository"
)

var repo = getRepo()

func getRepo() *repository.OtherRepository {
	return repository.NewOtherRepository()
}

func Get() ([]model.Other, error) {
	ret, err := repo.Get()
	return ret, err
}

func Create(Other *model.Other) (*model.Other, error) {
	ret, err := repo.Create(Other)
	return ret, err
}

func PutById(id uint, Other *model.Other) (*model.Other, error) {
	ret, err := repo.UpdateById(id, Other)
	return ret, err
}

func DeleteById(id uint) error {
	err := repo.DeleteById(id)
	return err
}

func GetById(id uint) (*model.Other, error) {
	ret, err := repo.GetById(id)
	return ret, err
}
