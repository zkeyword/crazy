package cmsServices

import (
	"CRAZY/model"
	"CRAZY/repository"
)

// CategoryService user服务
type CategoryService interface {
	Get(page int, pageSize int, title string) (*ReturnCategoryList, error)
	Create(Category *model.Category) (*model.Category, error)
	DeleteById(id int64) error
	PutCategoryById(id int64, Category *model.Category) (*model.Category, error)
	GetById(id int64) (*model.Category, error)
}

type categoryService struct {
	repo *repository.CategoryRepository
}

type ReturnCategoryList struct {
	Page     int              `json:"page"`
	PageSize int              `json:"pageSize"`
	Total    int64            `json:"total"`
	List     []model.Category `json:"list"`
}

// NewArticleService 实例化ArticleService
var NewCategoryService = newCategoryService()

func newCategoryService() CategoryService {
	return &categoryService{
		repo: repository.NewCategoryRepository(),
	}
}

func (s *categoryService) Get(page int, pageSize int, title string) (*ReturnCategoryList, error) {
	ret, err := s.repo.Get(page, pageSize, title)
	count := s.repo.GetCategoryCount(title)
	returnValue := &ReturnCategoryList{
		Page:     page,
		PageSize: pageSize,
		Total:    count,
		List:     ret,
	}
	return returnValue, err
}

func (s *categoryService) Create(Category *model.Category) (*model.Category, error) {
	ret, err := s.repo.Create(Category)
	return ret, err
}

func (s *categoryService) PutCategoryById(id int64, Category *model.Category) (*model.Category, error) {
	ret, err := s.repo.UpdateById(id, Category)
	return ret, err
}

func (s *categoryService) DeleteById(id int64) error {
	err := s.repo.DeleteById(id)
	return err
}

func (s *categoryService) GetById(id int64) (*model.Category, error) {
	ret, err := s.repo.GetById(id)
	return ret, err
}
