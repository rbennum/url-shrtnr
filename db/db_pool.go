package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Pool struct {
	instance *sqlx.DB
	statements map[string]*sqlx.Stmt
}

type PoolOptions struct {
	Host string
	Port int
	User string
	DBName string
	Pass string
}

var Pool_DB Pool

func Init(opts PoolOptions) error {
	psqlSetup := fmt.Sprintf(
		"host=%s port=%d user=%s dbname=%s sslmode=disable",
		opts.Host, 
		opts.Port, 
		opts.User, 
		opts.DBName,
	)
	if opts.Pass != "" {
		psqlSetup = fmt.Sprintf("%s password=%s", psqlSetup, opts.Pass)
	}
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
	statements := make(map[string]*sqlx.Stmt)
	for name, query := range queries() {
		statements[name], err = instance.Preparex(query)
		if err != nil {
			return err
		}
	}
	Pool_DB = Pool {
		instance: instance,
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
	return map[string]string {}
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

func (p Pool) CreateRole(tx *sql.Tx, roleName, password string) error {
    log.Printf("Creating role %s", roleName)
    _, err := tx.Exec(fmt.Sprintf(
        "CREATE ROLE %s WITH LOGIN PASSWORD '%s'",
        roleName,
        password,
    ))
    if err != nil {
        log.Printf("Error creating role: %v", err)
    }
    return err
}

func (p Pool) CheckRoleExists(tx *sql.Tx, roleName string) (bool, error) {
    query := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM pg_roles WHERE rolname='%s')", roleName)
    log.Printf("Checking role existence with query: %s", query)
    var exists bool
    err := tx.QueryRow(query).Scan(&exists)
    if err != nil {
        log.Printf("Error checking role existence: %v", err)
        return false, err
    }
    log.Printf("Role existence check result: %v", exists)
    return exists, nil
}

func (p Pool) SecureCreateRole(roleName string, password string) {
	tx, err := p.instance.Begin()
    if err != nil {
        log.Fatalf("Failed to begin transaction: %v", err)
    }

    // Create the role within the transaction
    if err := p.CreateRole(tx, roleName, password); err != nil {
        tx.Rollback()
        log.Fatalf("Failed to create role: %v", err)
    }

	// Commit the transaction
    if err := tx.Commit(); err != nil {
        log.Fatalf("Failed to commit transaction: %v", err)
    }
}

func (p Pool) SecureCheckRole(roleName string) {
	// Wait briefly before checking the role existence
    time.Sleep(1 * time.Second)

    // Check if the role exists
    tx, err := p.instance.Begin()
    if err != nil {
        log.Fatalf("Failed to begin transaction for role check: %v", err)
    }

    exists, err := p.CheckRoleExists(tx, roleName)
    if err != nil {
        tx.Rollback()
        log.Fatalf("Failed to check role existence: %v", err)
    }

    if err := tx.Commit(); err != nil {
        log.Fatalf("Failed to commit transaction for role check: %v", err)
    }

    if exists {
        log.Printf("Role %s exists.\n", roleName)
    } else {
        log.Printf("Role %s does not exist.\n", roleName)
    }
}

func (p Pool) GetInstance() *sqlx.DB {
	return p.instance
}
