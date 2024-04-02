package response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func WriteJSON(w http.ResponseWriter, status int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(v)
}

func WriteResponse(w http.ResponseWriter, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return json.NewEncoder(w).Encode(v)
}

func WriteError(w http.ResponseWriter, err error) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	return json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})
}
