package repository

import (
	"CRAZY/model"
	"CRAZY/utils/db"
)

type OtherRepository struct {
}

// Other 类型
type Other struct {
	ID    uint   `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  uint   `json:"type"`
}

func NewOtherRepository() *OtherRepository {
	return &OtherRepository{}
}

func (r *OtherRepository) Create(t *model.Other) (*model.Other, error) {
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Create(t).Error
	return t, err
}

func (r *OtherRepository) DeleteById(id uint) error {
	_db, err := db.GetMysql()
	if err != nil {
		return err
	}
	if err := _db.Where("id = ?", id).Delete(Other{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *OtherRepository) UpdateById(id uint, t *model.Other) (*model.Other, error) {
	var ret = new(model.Other)
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Model(&ret).Where("id=?", id).Updates(t).Error
	if err == nil {
		ret.ID = id
	}
	return ret, err
}

func (r *OtherRepository) GetById(id uint) (*model.Other, error) {
	var ret = &model.Other{}
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.First(ret, "id = ?", id).Error
	return ret, err
}

func (r *OtherRepository) Get() ([]model.Other, error) {
	var Other []model.Other
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Find(&Other).Error
	return Other, err
}
