package repositories

import (
	"github.com/rbennum/url-shrtnr/db"
	"github.com/rbennum/url-shrtnr/models"
)

type ShortRepository interface {
	GetURLByString(rand_str string, url string) (*models.Link, error)
	GetStaticURL() (*string, error)
}

type ShortRepository_Impl struct {
	Pool *db.Pool
}

func NewShortRepository(pool *db.Pool) ShortRepository {
	return ShortRepository_Impl {Pool: pool}
}

func (s ShortRepository_Impl) GetURLByString(
	rand_str string,
	url string,
) (*models.Link, error) {
	stmt := s.Pool.GetStatement("CreateShortURL")
	var url_obj models.Link
	err := stmt.QueryRowx(
		url,
		rand_str,
	).StructScan(&url_obj)
	if err != nil {
		return nil, err
	}
	return &url_obj, nil
}

func (s ShortRepository_Impl) GetStaticURL() (*string, error) {
	stmt := s.Pool.GetStatement("GetStaticURL")
	var saved_url *string
	err := stmt.Get(&saved_url)
	if err != nil {
		return nil, err
	}
	return saved_url, nil
}
