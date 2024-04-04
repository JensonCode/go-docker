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

	if r.Method == "POST" {
		return HandleCreateUser(w, r)
	}
	if r.Method == "PUT" {
		return HandleChangeUser(w, r)
	}
	if r.Method == "DELETE" {
		return HandleDeteleUser(w, r)
	}

	return fmt.Errorf("request method is not allowed: %s", r.Method)
}
