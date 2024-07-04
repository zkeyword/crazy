package cmsServices

import (
	"CRAZY/model"
	"CRAZY/repository"
	"CRAZY/utils"
	"strings"
)

// BannerService user服务
type BannerService interface {
	Get(page int, pageSize int, title string, categoryIds string) (*ReturnBannerList, error)
	Create(Banner *model.Banner, categoryIds string) (*model.Banner, error)
	DeleteById(id int64) error
	PutBannerById(id int64, Banner *model.Banner, categoryIds string) (*model.Banner, error)
	GetById(id int64) (*model.Banner, error)
}

type bannerService struct {
	repo        *repository.BannerRepository
	relatedRepo *repository.BannerCategoryRepository
}

type ReturnBannerList struct {
	Page     int            `json:"page"`
	PageSize int            `json:"pageSize"`
	Total    int64          `json:"total"`
	List     []model.Banner `json:"list"`
}

// NewArticleService 实例化ArticleService
var NewBannerService = newBannerService()

func newBannerService() BannerService {
	return &bannerService{
		repo: repository.NewBannerRepository(),
	}
}

func (s *bannerService) Get(page int, pageSize int, title string, categoryIds string) (*ReturnBannerList, error) {
	ret, err := s.repo.Get(page, pageSize, title)
	count := s.repo.GetBannerCount(title)
	returnValue := &ReturnBannerList{
		Page:     page,
		PageSize: pageSize,
		Total:    count,
		List:     ret,
	}
	return returnValue, err
}

func (s *bannerService) Create(Banner *model.Banner, categoryIds string) (*model.Banner, error) {
	ret, err := s.repo.Create(Banner)
	if err == nil {
		ids := strings.Split(categoryIds, ",")
		for _, v := range ids {
			s.relatedRepo.Create(ret.ID, utils.StrToUInt(v))
		}
	}
	return ret, err
}

func (s *bannerService) PutBannerById(id int64, Banner *model.Banner, categoryIds string) (*model.Banner, error) {
	ret, err := s.repo.UpdateById(id, Banner)
	ids := strings.Split(categoryIds, ",")
	s.relatedRepo.DeleteByPostId(id)
	for _, v := range ids {
		s.relatedRepo.Create(ret.ID, utils.StrToUInt(v))
	}
	return ret, err
}

func (s *bannerService) DeleteById(id int64) error {
	err := s.repo.DeleteById(id)
	if err == nil {
		s.relatedRepo.DeleteByPostId(id)
	}
	return err
}

func (s *bannerService) GetById(id int64) (*model.Banner, error) {
	ret, err := s.repo.GetById(id)
	return ret, err
}
