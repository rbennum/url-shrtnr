package client

import "github.com/jmoiron/sqlx"

type Client struct {
	DB *sqlx.DB
}
