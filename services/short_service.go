package services

import (
	"fmt"
	"log"

	"github.com/rbennum/url-shrtnr/models"
	"github.com/rbennum/url-shrtnr/repositories"
	"github.com/rbennum/url-shrtnr/utils"
)

type ShortService interface {
	GetURLByString(url string) (*models.Link, error)
}

type shortService_Impl struct {
	Repo repositories.ShortRepository
}

func NewShortService(repo repositories.ShortRepository) ShortService {
	return &shortService_Impl { Repo: repo }
}

func (s *shortService_Impl) GetURLByString(
	url string,
) (*models.Link, error) {
	rand_str := utils.RandomString(5)
	log.Printf("Log[GetURLByString] Created str: %s", rand_str)
	url_obj, err := s.Repo.GetURLByString(rand_str, url)
	if err != nil {
		return nil, err
	}
	static_url, err := s.Repo.GetStaticURL()
	if err != nil {
		return nil, err
	}
	url_obj.Tag = fmt.Sprintf("%s/%s", *static_url, url_obj.Tag)
	log.Printf("Log[GetURLByString] Error: %v", err)
	return url_obj, err
}