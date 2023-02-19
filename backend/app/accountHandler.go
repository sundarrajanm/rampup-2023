package app

import (
	"backend/dto"
	"backend/service"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type AccountHandlers struct {
	service service.AccountService
}

func (h AccountHandlers) NewAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]

	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, err := h.service.NewAccount(request)
		if err != nil {
			writeResponse(w, err.Code, err.OnlyMessage())
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}

func (h AccountHandlers) MakeTransaction(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	accountId := vars["account_id"]
	customerId := vars["customer_id"]

	var request dto.TransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		request.AccountId = accountId

		account, appError := h.service.MakeTransaction(request)

		if appError != nil {
			writeResponse(w, appError.Code, appError.OnlyMessage())
		} else {
			writeResponse(w, http.StatusOK, account)
		}
	}
}
