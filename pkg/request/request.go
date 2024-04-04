package request

import (
	"encoding/json"
	"net/http"
)

func ReadJSON[T any](r *http.Request, req *T) (*T, error) {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		return nil, err
	}

	return req, nil
}
