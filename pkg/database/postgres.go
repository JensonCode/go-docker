package database

import (
	"database/sql"
	"log"

	"github.com/JensonCode/go-docker/config"

	_ "github.com/lib/pq"
)

type PostgresDB struct {
	DB *sql.DB
}

var Postgres = new(PostgresDB)

func InitPostgres() error {
	log.Println("------   Connecting to Postgres DB   ------")
	defer log.Println("------ Postgres DB has been connected ------")

	conf := config.GetPostgresConfig()

	var err error

	Postgres.DB, err = sql.Open("postgres", conf)
	if err != nil {
		return err
	}

	if err := Postgres.DB.Ping(); err != nil {
		return err
	}

	return nil
}

func (p *PostgresDB) Close() {
	p.DB.Close()
}
