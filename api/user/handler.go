package user

import (
	"fmt"
	"net/http"
)

type UserHandler struct {
}

func HandleGetAccount(w http.ResponseWriter, r *http.Request) error {
	fmt.Println("handle get account")

	return nil
}
