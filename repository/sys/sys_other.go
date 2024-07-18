package sys

import (
	"CRAZY/model/sys"
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

func (r *OtherRepository) Create(t *sys.SysOther) (*sys.SysOther, error) {
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Create(t).Error
	return t, err
}

func (r *OtherRepository) BatchCreate(t []*sys.SysOther) ([]*sys.SysOther, error) {
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

func (r *OtherRepository) UpdateById(id uint, t *sys.SysOther) (*sys.SysOther, error) {
	var ret = new(sys.SysOther)
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

func (r *OtherRepository) GetById(id uint) (*sys.SysOther, error) {
	var ret = &sys.SysOther{}
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.First(ret, "id = ?", id).Error
	return ret, err
}

func (r *OtherRepository) Get() ([]sys.SysOther, error) {
	var Other []sys.SysOther
	_db, err := db.GetMysql()
	if err != nil {
		return nil, err
	}
	err = _db.Find(&Other).Error
	return Other, err
}
