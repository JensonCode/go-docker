package user

import (
	"fmt"
	"net/http"

	"github.com/JensonCode/go-docker/pkg/response"
)

func HandleGetUser(w http.ResponseWriter, r *http.Request) error {
	u, err := UserService.GetUser()
	if err != nil {
		return response.WriteError(w, err)
	}

	return response.WriteResponse(w, u)

}

func HandleCreateUser(w http.Response, r *http.Request) error {
	fmt.Println("handle create account")

	return nil
}
