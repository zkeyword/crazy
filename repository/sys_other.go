package repository

import (
	"CRAZY/model"
	"CRAZY/utils/db"
	"fmt"
)

type OtherRepository struct {
}

// Other 类型
type Other struct {
	ID    uint
	Key   string
	Value string
	Type  uint
}

func NewOtherRepository() *OtherRepository {
	return &OtherRepository{}
}

func (r *OtherRepository) Create(t *model.Other) (*model.Other, error) {
	err := db.GetMysql().Create(t).Error
	return t, err
}

func (r *OtherRepository) DeleteById(id int64) error {
	fmt.Println(id)
	if err := db.GetMysql().Where("id = ?", id).Delete(Other{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *OtherRepository) UpdateById(id int64, t *model.Other) (*model.Other, error) {
	var ret = new(model.Other)
	err := db.GetMysql().Model(&ret).Where("id=?", id).Updates(t).Error
	return ret, err
}

func (r *OtherRepository) GetById(id int64) (*model.Other, error) {
	var ret = &model.Other{}
	err := db.GetMysql().First(ret, "id = ?", id).Error
	return ret, err
}

func (r *OtherRepository) Get() ([]model.Other, error) {
	var Other []model.Other
	err := db.GetMysql().Find(&Other).Error
	return Other, err
}
