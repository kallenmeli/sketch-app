package routers

import (
	"encoding/json"
	"net/http"
)

func ToJSON[T any](w http.ResponseWriter, statusCode int, body T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(body)
}

func FromJSON[T any](r *http.Request) (T, error) {
	var result T
	err := json.NewDecoder(r.Body).Decode(&result)
	return result, err
}

func NotFound(w http.ResponseWriter, body error) error {
	_ = ToJSON(w, http.StatusNotFound, body)
	return body
}
