package handler

import (
	"net/http"

	"github.com/JensonCode/go-docker/pkg/response"
)

type handler func(http.ResponseWriter, *http.Request) error

func HttpHandler(f handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			response.WriteError(w, err)
		}
	}
}
