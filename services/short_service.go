package services

import (
	"fmt"
	"log"
	"os"

	"github.com/rbennum/url-shrtnr/models"
	"github.com/rbennum/url-shrtnr/repositories"
	"github.com/rbennum/url-shrtnr/utils"
)

type ShortService interface {
	CreateURL(url string) (*models.Link, error)
	GetURL(tag string) (*string, error)
}

type shortService_Impl struct {
	Repo   repositories.ShortRepository
	Config utils.CommonConfig
}

func NewShortService(
	repo repositories.ShortRepository,
	config utils.CommonConfig,
) ShortService {
	return &shortService_Impl{Repo: repo, Config: config}
}

func (s *shortService_Impl) CreateURL(
	url string,
) (*models.Link, error) {
	rand_str := utils.RandomString(5)
	log.Printf("Log[GetURLByString] Created str: %s", rand_str)
	url_obj, err := s.Repo.CreateURL(rand_str, url)
	if err != nil {
		return nil, err
	}
	static_url := ""
	if os.Getenv("GIN_MODE") == "release" {
		static_url = s.Config.StaticURL
	} else {
		static_url = fmt.Sprintf(
			"%s:%s",
			s.Config.StaticURL,
			s.Config.ShortServerPort,
		)
	}
	url_obj.Tag = fmt.Sprintf("%s/%s", static_url, url_obj.Tag)
	log.Printf("Log[GetURLByString] Error: %v", err)
	return url_obj, err
}

func (s *shortService_Impl) GetURL(tag string) (*string, error) {
	url, err := s.Repo.GetURL(tag)
	if err != nil {
		return nil, err
	}
	return url, nil
}
