package cmsServices

import (
	"CRAZY/model"
	"CRAZY/repository"
	"CRAZY/utils"
	"strings"
)

// PostService user服务
type PostService interface {
	Get(page int, pageSize int, title string, categoryIds string) (*ReturnPostList, error)
	GetNews() (*ReturnPostList, error)
	Create(Post *model.Post, categoryIds string) (*model.Post, error)
	DeleteById(id int64) error
	PutPostById(id int64, Post *model.Post, categoryIds string) (*model.Post, error)
	GetById(id int64) (*model.Post, error)
	PostNewsSet(action string, ids []int64) error
}

type postService struct {
	repo        *repository.PostRepository
	relatedRepo *repository.PostCategoryRepository
}

type ReturnPostList struct {
	Page     int          `json:"page"`
	PageSize int          `json:"pageSize"`
	Total    int64        `json:"total"`
	List     []model.Post `json:"list"`
}

// NewArticleService 实例化ArticleService
var NewPostService = newPostService()

func newPostService() PostService {
	return &postService{
		repo: repository.NewPostRepository(),
	}
}

func (s *postService) Get(page int, pageSize int, title string, categoryIds string) (*ReturnPostList, error) {
	ret, err := s.repo.Get(page, pageSize, title, categoryIds)
	count := s.repo.GetPostCount(title)
	returnValue := &ReturnPostList{
		Page:     page,
		PageSize: pageSize,
		Total:    count,
		List:     ret,
	}
	return returnValue, err
}

func (s *postService) Create(Post *model.Post, categoryIds string) (*model.Post, error) {
	ret, err := s.repo.Create(Post)
	if err == nil {
		ids := strings.Split(categoryIds, ",")
		for _, v := range ids {
			s.relatedRepo.Create(ret.ID, utils.StrToUInt(v))
		}
	}
	return ret, err
}

func (s *postService) PutPostById(id int64, Post *model.Post, categoryIds string) (*model.Post, error) {
	ret, err := s.repo.UpdateById(id, Post)
	ids := strings.Split(categoryIds, ",")
	s.relatedRepo.DeleteByPostId(id)
	for _, v := range ids {
		s.relatedRepo.Create(ret.ID, utils.StrToUInt(v))
	}
	return ret, err
}

func (s *postService) DeleteById(id int64) error {
	err := s.repo.DeleteById(id)
	if err == nil {
		s.relatedRepo.DeleteByPostId(id)
	}
	return err
}

func (s *postService) GetById(id int64) (*model.Post, error) {
	ret, err := s.repo.GetById(id)
	return ret, err
}

func (s *postService) GetNews() (*ReturnPostList, error) {
	ret, err := s.repo.GetNews()
	returnValue := &ReturnPostList{
		Page:     1,
		PageSize: len(ret),
		Total:    int64(len(ret)),
		List:     ret,
	}
	return returnValue, err
}

func (s *postService) PostNewsSet(action string, ids []int64) error {
	return s.repo.PostNewsSet(action, ids)
}
