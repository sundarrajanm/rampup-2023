//	Comapany RampUp2023:
//	 version: 0.0.1
//	 title: Ramp Up and Learn in 2023
//	Schemes: http, https
//	Host: localhost:8000
//	BasePath: /
//	Produces:
//	  - application/json
//
// swagger:meta
package controller

import (
	"banking-resource-api/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type CustomerHandler struct {
	Service service.CustomerService
}

// swagger:route GET /customers
// Get customers list
//
// responses:
//
//	500: errs.AppError
//	200: []dto.CustomerResponse
func (ch CustomerHandler) GetAllCustomers(rw http.ResponseWriter, r *http.Request) {
	customers, appError := ch.Service.GetAllCustomers()

	if appError != nil {
		writeResponse(rw, appError.Code, appError)
	} else {
		writeResponse(rw, http.StatusOK, customers)
	}
}

// swagger:route  GET /customers/{id}
// Get a customer
//
// responses:
//
//	404: errs.AppError
//	500: errs.AppError
//	200: dto.CustomerResponse
func (ch CustomerHandler) GetCustomerById(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["customer_id"]

	customer, appError := ch.Service.GetCustomerById(id)
	if appError != nil {
		writeResponse(rw, appError.Code, appError)
	} else {
		writeResponse(rw, http.StatusOK, customer)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func NewCustomerHandler(service service.CustomerService) CustomerHandler {
	return CustomerHandler{Service: service}
}
