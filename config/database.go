package config

import (
	"fmt"
	"os"
)

type PostgresConfig struct {
	Host     string
	Port     string
	User     string
	Dbname   string
	Password string
	Sslmode  string
}

func GetPostgresConfig() string {

	var conf = new(PostgresConfig)

	conf.Host = os.Getenv("POSTGRES_HOST")
	conf.Port = os.Getenv("POSTGRES_PORT")
	conf.User = os.Getenv("POSTGRES_USER")
	conf.Dbname = os.Getenv("POSTGRES_DB")
	conf.Password = os.Getenv("POSTGRES_PASSWORD")
	conf.Sslmode = os.Getenv("POSTGRES_SSLMODE")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Dbname,
		conf.Password,
		conf.Sslmode,
	)

	return connectionString

}
