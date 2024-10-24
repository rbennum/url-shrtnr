package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/rbennum/url-shrtnr/internals/models"
)

type UrlRepo struct {
	db *sqlx.DB
}

func NewUrlRepo(db *sqlx.DB) *UrlRepo {
	return &UrlRepo{db: db}
}

// Define CRUD operations here
func (r *UrlRepo) CreateUrl(originalUrl string, shortTag string) (*models.Link, error) {
	// Prepare the SQL statement
	stmt, err := r.db.Preparex(`
		INSERT INTO link_mappers (url, short_tag)
		VALUES ($1, $2)
		RETURNING id, url, short_tag, created_at
	`)
	if err != nil {
		return nil, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close() // Ensure the prepared statement is closed after execution

	var link models.Link
	// Execute the prepared statement with the provided URL and short tag
	err = stmt.QueryRowx(originalUrl, shortTag).StructScan(&link)
	if err != nil {
		return nil, fmt.Errorf("failed to execute prepared statement: %w", err)
	}

	return &link, nil
}
