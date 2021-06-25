package services

import (
	"CRAZY/model"
	"CRAZY/repository"
)

// OtherService user服务
type OtherService interface {
	Create(Other *model.Other) (*model.Other, error)
	Get() ([]model.Other, error)
	GetById(id int64) (*model.Other, error)
	DeleteById(id int64) error
	PutById(id int64, Other *model.Other) (*model.Other, error)
}

type otherRepository struct {
	repo *repository.OtherRepository
}

// NewArticleService 实例化ArticleService
var NewOtherService = newOtherService()

func newOtherService() OtherService {
	return &otherRepository{
		repo: repository.NewOtherRepository(),
	}
}

func (s *otherRepository) Get() ([]model.Other, error) {
	ret, err := s.repo.Get()
	return ret, err
}

func (s *otherRepository) Create(Other *model.Other) (*model.Other, error) {
	ret, err := s.repo.Create(Other)
	return ret, err
}

func (s *otherRepository) PutById(id int64, Other *model.Other) (*model.Other, error) {
	ret, err := s.repo.UpdateById(id, Other)
	return ret, err
}

func (s *otherRepository) DeleteById(id int64) error {
	err := s.repo.DeleteById(id)
	return err
}

func (s *otherRepository) GetById(id int64) (*model.Other, error) {
	ret, err := s.repo.GetById(id)
	return ret, err
}
