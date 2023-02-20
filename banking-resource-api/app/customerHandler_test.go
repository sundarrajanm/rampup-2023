package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func Test_should_return_empty_customers_with_status_code_200(t *testing.T) {
	// customers := []dto.CustomerResponse{}

	router := mux.NewRouter()
	ch := CustomerHandler{}
	router.HandleFunc("/customers", ch.GetAllCustomers)

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	responseWriter := httptest.NewRecorder()
	router.ServeHTTP(responseWriter, request)

	if responseWriter.Code != http.StatusOK {
		t.Errorf("Got response code: %d", responseWriter.Code)
	}
}
