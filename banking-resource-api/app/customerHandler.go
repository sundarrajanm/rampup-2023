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
	rw.WriteHeader(200)
	rw.Header().Add("Content-Type", "application/json")

	customers, appError := ch.Service.GetAllCustomers()

	if appError != nil {
		writeResponse(rw, appError.Code, appError)
	} else {
		writeResponse(rw, http.StatusOK, customers)
	}
}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	if code != http.StatusOK {
		w.WriteHeader(code)
	}
	json.NewEncoder(w).Encode(data)
}
