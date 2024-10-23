package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/rbennum/url-shrtnr/internals/models"
	mb "github.com/rbennum/url-shrtnr/internals/rabbitmq"
	"github.com/rbennum/url-shrtnr/internals/repository"
	"github.com/rbennum/url-shrtnr/utils"
	"github.com/rs/zerolog/log"
)

type UrlService interface {
	CreateUrl(ctx context.Context, url string) (*models.Link, error)
}

type urlService_Impl struct {
	repo   *repository.UrlRepo
	config *utils.CommonConfig
}

func NewUrlService(repo *repository.UrlRepo, config *utils.CommonConfig) UrlService {
	return &urlService_Impl{repo: repo, config: config}
}

func (s *urlService_Impl) CreateUrl(ctx context.Context, originalUrl string) (*models.Link, error) {
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

	msg, _ := json.Marshal(models.URLData{URL: originalUrl, ShortTag: shortTag})
	log.Debug().
		Msgf("%s %s", originalUrl, shortTag)
	mb.SendMessage(msg, ctx)

	return link, nil
}
