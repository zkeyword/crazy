package cmsServices

import (
	"CRAZY/model/cms"
	repository "CRAZY/repository/cms"
)

// CategoryService user服务
type CategoryService interface {
	Get(page int, pageSize int, title string) (*ReturnCategoryList, error)
	Create(Category *cms.CmsCategory) (*cms.CmsCategory, error)
	DeleteById(id int64) error
	PutCategoryById(id int64, Category *cms.CmsCategory) (*cms.CmsCategory, error)
	GetById(id int64) (*cms.CmsCategory, error)
}

type categoryService struct {
	repo *repository.CategoryRepository
}

type ReturnCategoryList struct {
	Page     int               `json:"page"`
	PageSize int               `json:"pageSize"`
	Total    int64             `json:"total"`
	List     []cms.CmsCategory `json:"list"`
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

func (s *categoryService) Create(Category *cms.CmsCategory) (*cms.CmsCategory, error) {
	ret, err := s.repo.Create(Category)
	return ret, err
}

func (s *categoryService) PutCategoryById(id int64, Category *cms.CmsCategory) (*cms.CmsCategory, error) {
	ret, err := s.repo.UpdateById(id, Category)
	return ret, err
}

func (s *categoryService) DeleteById(id int64) error {
	err := s.repo.DeleteById(id)
	return err
}

func (s *categoryService) GetById(id int64) (*cms.CmsCategory, error) {
	ret, err := s.repo.GetById(id)
	return ret, err
}
