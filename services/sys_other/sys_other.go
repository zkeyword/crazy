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

func PutById(id int64, Other *model.Other) (*model.Other, error) {
	ret, err := repo.UpdateById(id, Other)
	return ret, err
}

func DeleteById(id int64) error {
	err := repo.DeleteById(id)
	return err
}

func GetById(id int64) (*model.Other, error) {
	ret, err := repo.GetById(id)
	return ret, err
}
