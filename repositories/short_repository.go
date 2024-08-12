package repositories

import (
	"github.com/rbennum/url-shrtnr/db"
	"github.com/rbennum/url-shrtnr/models"
)

type ShortRepository interface {
	GetURLByString(str string) (*models.Link)
}

type ShortRepository_Impl struct {
	Pool *db.Pool
}

func NewShortRepository(pool *db.Pool) ShortRepository {
	return ShortRepository_Impl {Pool: pool}
}

func (s ShortRepository_Impl) GetURLByString(str string) (*models.Link) {
	return &models.Link{}
}