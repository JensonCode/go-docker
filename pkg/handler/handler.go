package handler

import (
	"net/http"

	"github.com/JensonCode/go-docker/pkg/response"
)

type Services func(http.ResponseWriter, *http.Request) error

func HttpHandler(f Services) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			response.WriteError(w, err)
		}
	}
}
