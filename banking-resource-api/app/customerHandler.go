package app

import (
	"banking-resource-api/service"
	"encoding/json"
	"net/http"
)

type CustomerHandler struct {
	Service service.CustomerService
}

func (ch CustomerHandler) GetAllCustomers(rw http.ResponseWriter, r *http.Request) {
	customers, appError := ch.Service.GetAllCustomers()

	if appError != nil {
		writeResponse(rw, appError.Code, appError)
	} else {
		writeResponse(rw, http.StatusOK, customers)
	}
}

func (ch CustomerHandler) GetCustomerById(rw http.ResponseWriter, r *http.Request) {

}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func NewCustomerHandler(service service.CustomerService) CustomerHandler {
	return CustomerHandler{Service: service}
}
