package cmsServices

import (
	"CRAZY/model/cms"
	repository "CRAZY/repository/cms"
)

// BannerService user服务
type BannerService interface {
	Get(page int, pageSize int, title string) (*ReturnBannerList, error)
	Create(Banner *cms.CmsBanner) (*cms.CmsBanner, error)
	DeleteById(id int64) error
	PutBannerById(id int64, Banner *cms.CmsBanner) (*cms.CmsBanner, error)
	GetById(id int64) (*cms.CmsBanner, error)
}

type bannerService struct {
	repo *repository.BannerRepository
}

type ReturnBannerList struct {
	Page     int             `json:"page"`
	PageSize int             `json:"pageSize"`
	Total    int64           `json:"total"`
	List     []cms.CmsBanner `json:"list"`
}

// NewArticleService 实例化ArticleService
var NewBannerService = newBannerService()

func newBannerService() BannerService {
	return &bannerService{
		repo: repository.NewBannerRepository(),
	}
}

func (s *bannerService) Get(page int, pageSize int, title string) (*ReturnBannerList, error) {
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

func (s *bannerService) Create(Banner *cms.CmsBanner) (*cms.CmsBanner, error) {
	ret, err := s.repo.Create(Banner)
	return ret, err
}

func (s *bannerService) PutBannerById(id int64, Banner *cms.CmsBanner) (*cms.CmsBanner, error) {
	ret, err := s.repo.UpdateById(id, Banner)
	return ret, err
}

func (s *bannerService) DeleteById(id int64) error {
	err := s.repo.DeleteById(id)
	return err
}

func (s *bannerService) GetById(id int64) (*cms.CmsBanner, error) {
	ret, err := s.repo.GetById(id)
	return ret, err
}
