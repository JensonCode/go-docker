package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/JensonCode/go-docker/configs"
	_ "github.com/lib/pq"
)

type PostgresDB struct {
	db *sql.DB
}

func NewDBInstance() (*PostgresDB, error) {
	log.Println("----- Connecting to Postgres DB -----")

	conf := configs.GetPostgresConfig()

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s",
		conf.User,
		conf.Dbname,
		conf.Password,
		conf.Sslmode)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if db.Ping() != nil {
		return nil, err
	}

	log.Println("----- Postgres DB has been connected -----")

	return &PostgresDB{db: db}, nil
}

func (p *PostgresDB) Init() error {

	//todo: create table and default rows

	return nil
}
