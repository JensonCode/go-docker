package main

import (
	"log"
	"net/http"

	server "github.com/JensonCode/go-docker/api"
	"github.com/JensonCode/go-docker/database"
)

func main() {

	db, err := database.NewDBInstance()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Init(); err != nil {
		log.Fatal(err)
	}

	apiServer := server.Run(db)

	apiServer.RegisterRouters()

	http.ListenAndServe(apiServer.Port, apiServer.Router)
}
