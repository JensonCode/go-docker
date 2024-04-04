package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JensonCode/go-docker/pkg/handler"

	"github.com/gorilla/mux"
)

type UserRouter struct{}

var Router = new(UserRouter)

func (router *UserRouter) Register(r *mux.Router) {
	log.Println("Resiger User Router")

	r.HandleFunc("/user", handler.HttpHandler(userHanlders))
}

func userHanlders(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return HandleGetUser(w, r)
	}
	if r.Method == "POST" {
		return HandleGetUser(w, r)
	}
	if r.Method == "DELETE" {
		return HandleGetUser(w, r)
	}

	return fmt.Errorf("request method is not allowed: %s", r.Method)
}
