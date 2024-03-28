package configs

import (
	"os"
)

type APIServer struct {
	Port string
	DB   Database
}

func NewAPIServer(db Database) *APIServer {
	return &APIServer{
		Port: os.Getenv("PORT"),
		DB:   db,
	}
}
