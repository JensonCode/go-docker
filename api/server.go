package server

import (
	"log"

	"github.com/JensonCode/go-docker/api/user"
	"github.com/JensonCode/go-docker/configs"
	"github.com/JensonCode/go-docker/database"
	"github.com/gorilla/mux"
)

type APIServer struct {
	Router *mux.Router
	Port   string
	DB     *database.PostgresDB
}

func Run(db *database.PostgresDB) *APIServer {
	log.Println("***** Runs API server *****")

	router := mux.NewRouter()

	conf := configs.GetServerConfig()

	return &APIServer{
		Router: router,
		Port:   conf.Port,
		DB:     db,
	}

}

func (s *APIServer) RegisterRouters() {
	log.Println("----- Register Routers----- ")

	user.Router.Register(s.Router)

	log.Println("----- Register Success----- ")

}
