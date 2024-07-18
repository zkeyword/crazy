package otherService

import (
	"CRAZY/model/sys"
	repository "CRAZY/repository/sys"
)

var repo = getRepo()

func getRepo() *repository.OtherRepository {
	return repository.NewOtherRepository()
}

func Get() ([]sys.SysOther, error) {
	ret, err := repo.Get()
	return ret, err
}

func Create(Other *sys.SysOther) (*sys.SysOther, error) {
	ret, err := repo.Create(Other)
	return ret, err
}

func BatchCreate(Other []*sys.SysOther) ([]*sys.SysOther, error) {
	ret, err := repo.BatchCreate(Other)
	return ret, err
}

func PutById(id uint, Other *sys.SysOther) (*sys.SysOther, error) {
	ret, err := repo.UpdateById(id, Other)
	return ret, err
}

func DeleteById(id uint) error {
	err := repo.DeleteById(id)
	return err
}

func GetById(id uint) (*sys.SysOther, error) {
	ret, err := repo.GetById(id)
	return ret, err
}
