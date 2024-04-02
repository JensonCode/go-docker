package configs

import "os"

type ServerConfig struct {
	Port string
}

func GetServerConfig() *ServerConfig {
	return &ServerConfig{
		Port: os.Getenv("PORT"),
	}
}

type PostgresConfig struct {
	User     string
	Dbname   string
	Password string
	Sslmode  string
}

func GetPostgresConfig() *PostgresConfig {
	return &PostgresConfig{
		User:     os.Getenv("POSTGRES_USER"),
		Dbname:   os.Getenv("POSTGRES_DB"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Sslmode:  os.Getenv("POSTGRES_SSLMODE"),
	}
}
