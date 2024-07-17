package system

import (
	"CRAZY/model/system"
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

func (r *OtherRepository) Create(t *system.Other) (*system.Other, error) {
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Create(t).Error
	return t, err
}

func (r *OtherRepository) BatchCreate(t []*system.Other) ([]*system.Other, error) {
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
	return _db.Where("id = ?", id).Delete(Other{}).Error
}

func (r *OtherRepository) UpdateById(id uint, t *system.Other) (*system.Other, error) {
	var ret = new(system.Other)
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

func (r *OtherRepository) GetById(id uint) (*system.Other, error) {
	var ret = &system.Other{}
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.First(ret, "id = ?", id).Error
	return ret, err
}

func (r *OtherRepository) Get() ([]system.Other, error) {
	var Other []system.Other
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Find(&Other).Error
	return Other, err
}
