package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JensonCode/go-docker/internal/configs"
	"github.com/gorilla/mux"
)

func main() {

	db, err := configs.NewPostgresInstance()
	if err != nil {
		log.Fatal(err)
	}

	//todo: init job
	// if err := db.Init(); err != nil {
	// 	log.Fatal(err)
	// }

	server := configs.NewAPIServer(db)

	router := mux.NewRouter()

	//router.register

	fmt.Println("Server listening")

	http.ListenAndServe(server.Port, router)

}
