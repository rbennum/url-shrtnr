package service

import (
	"fmt"

	"github.com/rbennum/url-shrtnr/internals/models"
	"github.com/rbennum/url-shrtnr/internals/repository"
	"github.com/rbennum/url-shrtnr/utils"
)

type UrlService interface {
	CreateUrl(url string) (*models.Link, error)
}

type urlService_Impl struct {
	repo   *repository.UrlRepo
	config *utils.CommonConfig
}

func NewUrlService(repo *repository.UrlRepo, config *utils.CommonConfig) UrlService {
	return &urlService_Impl{repo: repo, config: config}
}

func (s *urlService_Impl) CreateUrl(originalUrl string) (*models.Link, error) {
	// Step 1: Generate a 5-character random short_tag
	shortTag := utils.RandomString(5)

	// Step 2: Save the URL with the generated short_tag to the repository
	link, err := s.repo.CreateUrl(originalUrl, shortTag)
	if err != nil {
		return nil, err
	}

	// Step 3: Concatenate base URL with short_tag
	baseUrl := s.config.StaticShortURL
	link.Tag = fmt.Sprintf("%s/%s", baseUrl, link.Tag)

	return link, nil
}
