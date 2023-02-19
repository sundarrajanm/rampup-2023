package app

import (
	"backend/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandlers struct {
	service service.CustomerService
}

func (ch *CustomerHandlers) getAllCustomers(rw http.ResponseWriter, r *http.Request) {

	customers, err := ch.service.GetAllCustomers()

	if err != nil {
		writeResponse(rw, err.Code, err.OnlyMessage())
	} else {
		writeResponse(rw, http.StatusOK, customers)
	}
}

func (ch *CustomerHandlers) getCustomer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, err := ch.service.GetCustomer(id)
	if err != nil {
		writeResponse(rw, err.Code, err.OnlyMessage())
	} else {
		writeResponse(rw, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		panic(err)
	}
}
