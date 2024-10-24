package db

import "github.com/jmoiron/sqlx"

func NewDatabase(databaseUrl string) *sqlx.DB {
	db, err := sqlx.Open("postgres", databaseUrl)
	if err != nil {
		panic(err)
	}
	// for good luck
	_, err = db.Query("SELECT 2+2")
	if err != nil {
		panic(err)
	}
	return db
}
