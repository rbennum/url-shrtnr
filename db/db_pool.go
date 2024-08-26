package db

import (
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/rbennum/url-shrtnr/utils"
)

type Pool struct {
	instance   *sqlx.DB
	statements map[string]*sqlx.Stmt
}

var Pool_DB *Pool

func Init(config *utils.CommonConfig) error {
	// init DB
	psqlSetup := config.DBSourceName
	var err error
	instance, err := sqlx.Connect("postgres", psqlSetup)
	if err != nil {
		return err
	}
	// activating client inside of the pool
	_, err = instance.Query("SELECT 2 + 2;")
	if err != nil {
		return err
	}
	// migrate necessary schemas
	migrateDB(instance.DB, config.DBName)
	// generate statements
	statements := make(map[string]*sqlx.Stmt)
	for name, query := range queries() {
		statements[name], err = instance.Preparex(query)
		if err != nil {
			return err
		}
	}
	Pool_DB = &Pool{
		instance:   instance,
		statements: statements,
	}
	return nil
}

func queries() map[string]string {
	/*
		Example:
		"GetAllUsers": "SELECT * FROM users;",
		"GetUserByID": "SELECT * FROM users WHERE id=$1;",
		"CreateUser": `
			INSERT INTO users (username, bio)
			VALUES ($1, $2)
			RETURNING *;
		`,
		"UpdateUser": `
			UPDATE users
			SET username = COALESCE($1, username), bio = COALESCE($2, bio)
			WHERE id = $3
			RETURNING *;
		`,
		"DeleteUser": `
			DELETE FROM users
			WHERE id = $1;
		`,
		"CountUsers": "SELECT COUNT(*) FROM users;"
	*/
	return map[string]string{
		"CreateShortURL": `
			INSERT INTO link_mappers (url, short_tag)
			VALUES ($1, $2)
			RETURNING *;
		`,
		"GetStaticURL": `
			SELECT url FROM static_url;
		`,
		"GetURL": `
			SELECT url
			FROM link_mappers
			WHERE short_tag = $1;
		`,
	}
}

func migrateDB(instance *sql.DB, dbName string) {
	driver, err := postgres.WithInstance(
		instance,
		&postgres.Config{},
	)
	if err != nil {
		panic(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://db/migrations",
		dbName,
		driver,
	)
	if err != nil {
		panic(err)
	}
	vers, is_dirty, _ := m.Version()
	log.Printf("Migrating check up: %d, %t", vers, is_dirty)
	if is_dirty {
		m.Force(int(vers) - 1)
	}
	err = m.Up()
	if err != nil {
		if err != migrate.ErrNoChange {
			panic(err)
		}
	}
}

func (p Pool) GetStatement(name string) *sqlx.Stmt {
	return p.statements[name]
}

func (p Pool) Close() {
	for _, stmt := range p.statements {
		stmt.Close()
	}
	p.instance.Close()
}

func (p Pool) Query(q string) (*sqlx.Rows, error) {
	rows, err := p.instance.Queryx(q)
	return rows, err
}

func (p Pool) QueryRow(q string, dest ...interface{}) error {
	return p.instance.QueryRow(q).Scan(dest...)
}

func (p Pool) Exec(q string) error {
	tx, err := p.instance.Begin()
	if err != nil {
		log.Fatalf("Failed to begin a transaction %v", err)
	}
	_, err = p.instance.Exec(q)
	if err != nil {
		tx.Rollback()
		log.Fatalf("Failed to create a role %v", err)
	}
	return tx.Commit()
}

func (p Pool) GetInstance() *sqlx.DB {
	return p.instance
}
