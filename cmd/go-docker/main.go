package main

import (
	"log"
	"net/http"

	"github.com/JensonCode/go-docker/internal/api/user"
	"github.com/JensonCode/go-docker/internal/models"
	"github.com/JensonCode/go-docker/pkg/database"
	"github.com/JensonCode/go-docker/pkg/server"
	"github.com/gorilla/mux"
)

func main() {

	err := database.InitPostgres()

	db := database.Postgres
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	if err := migrate(); err != nil {
		log.Fatal(err)
	}

	apiServer := server.Run()

	registerRouters(apiServer.Router)

	http.ListenAndServe(apiServer.Port, apiServer.Router)
}

func registerRouters(r *mux.Router) {
	log.Println("----- Register Routers----- ")

	user.Router.Register(r)
	log.Println("----- Register Success----- ")

}

func migrate() error {
	if err := createUserTable(); err != nil {
		return err
	}
	if err := createDefaultUser(); err != nil {
		return err
	}

	return nil
}

func createUserTable() error {
	log.Println("creating user table")

	query := `create table if not exists users (
		id serial primary key,
		username varchar(100),
		password varchar(100),
		created_at timestamp
	)`

	_, err := database.Postgres.DB.Exec(query)

	return err
}

func createDefaultUser() error {
	log.Println("creating default user")
	defer log.Println("Complete create default user")
	isCreated, err := user.UserService.IsUsernameExist("admin")
	if err != nil {
		return err
	}

	if isCreated {
		return nil
	}

	_, err = user.UserService.Create(
		&models.CreateUserRequest{
			Username: "admin",
			Password: "admin",
		},
	)

	if err != nil {
		return err
	}

	return nil
}
