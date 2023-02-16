package app

import (
	"backend/service"
	"encoding/json"
	"net/http"
)

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(rw http.ResponseWriter, r *http.Request) {

	customers, _ := ch.service.GetAllCustomers()

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(customers)
}
