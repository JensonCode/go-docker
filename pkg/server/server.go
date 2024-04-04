package server

import (
	"log"
	"os"

	"github.com/gorilla/mux"
)

type APIServer struct {
	Router *mux.Router
	Port   string
}

func Run() *APIServer {
	log.Println("***** Runs API server *****")

	router := mux.NewRouter()

	return &APIServer{
		Router: router,
		Port:   os.Getenv("PORT"),
	}
}
