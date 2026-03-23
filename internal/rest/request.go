package rest

import (
	"encoding/json"
	"log"
	"net/http"
)

func Decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		log.Println(err)
		return v, err
	}

	return v, nil
}
