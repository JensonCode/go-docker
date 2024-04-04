package user

import (
	"fmt"
	"net/http"

	"github.com/JensonCode/go-docker/internal/models"
	"github.com/JensonCode/go-docker/pkg/request"
	"github.com/JensonCode/go-docker/pkg/response"
)

func HandleCreateUser(w http.ResponseWriter, r *http.Request) error {
	req := new(models.CreateUserRequest)
	req, err := request.ReadJSON(r, req)
	if err != nil {
		return response.WriteError(w, err)
	}

	u, err := UserService.Create(req)
	if err != nil {
		return response.WriteError(w, err)
	}

	return response.WriteResponse(w, u)
}

func HandleUpdateUser(w http.ResponseWriter, r *http.Request) error {

	req := new(models.UpdateUserRequest)
	req, err := request.ReadJSON(r, req)
	if err != nil {
		return response.WriteError(w, err)
	}

	u, err := UserService.Update(req)
	if err != nil {
		return response.WriteError(w, err)
	}

	return response.WriteResponse(w, u)
}

func HandleDeteleUser(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("handle delete account")

	return nil
}
