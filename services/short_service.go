package services

import (
	"github.com/rbennum/url-shrtnr/models"
	"github.com/rbennum/url-shrtnr/repositories"
)

type ShortService interface {
	GetURLByString(str string) (*models.Link, error)
}

type ShortService_Impl struct {
	Repo *repositories.ShortRepository
}

func NewShortService(repo *repositories.ShortRepository) ShortService {
	return &ShortService_Impl { Repo: repo }
}

func (s ShortService_Impl) GetURLByString(
	str string,
) (*models.Link, error) {
	return nil, nil
}