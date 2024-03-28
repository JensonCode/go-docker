package configs

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	user     = os.Getenv("POSTGRES_USER")
	dbname   = os.Getenv("POSTGRES_DBNAME")
	password = os.Getenv("POSTGRES_PASSWORD")
	sslmode  = os.Getenv("POSTGRES_SSLMODE")
)

type Database interface {
}

type PostgresDB struct {
	db *sql.DB
}

func NewPostgresInstance() (*PostgresDB, error) {

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", user, dbname, password, sslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if db.Ping() != nil {
		return nil, err
	}

	return &PostgresDB{
		db: db,
	}, nil
}
