package app

import "net/http"

type CustomerHandler struct {
}

func (ch CustomerHandler) GetAllCustomers(rw http.ResponseWriter, r *http.Request) {
	rw.WriteHeader(200)
}
