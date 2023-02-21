package app

import (
	"banking-resource-api/service"
	"encoding/json"
	"net/http"
)

type CustomerHandler struct {
	service service.CustomerService
}

func (ch CustomerHandler) GetAllCustomers(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(200)
	rw.Header().Add("Content-Type", "application/json")

	encoder := json.NewEncoder(rw)
	encoder.Encode(ch.service.GetAllCustomers())
}
