package app

import (
	"encoding/json"
	"net/http"
)

type CustomerHandler struct {
}

func (ch CustomerHandler) GetAllCustomers(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(200)
	rw.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(rw)
	encoder.Encode([]any{})
}
