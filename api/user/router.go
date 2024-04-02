package user

import (
	"fmt"
	"log"
	"net/http"

	"github.com/JensonCode/go-docker/pkg/handler"
	"github.com/gorilla/mux"
)

type UserRouter struct {
}

var Router = new(UserRouter)

func (r *UserRouter) Register(router *mux.Router) {
	log.Println("Resiger User Router")
	router.HandleFunc("/user", handler.HttpHandler(UserServices))
}

func UserServices(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {
		return HandleGetAccount(w, r)
	}
	if r.Method == "POST" {
		return HandleGetAccount(w, r)
	}
	if r.Method == "DELETE" {
		return HandleGetAccount(w, r)
	}

	return fmt.Errorf("request method is not allowed: %s", r.Method)
}
