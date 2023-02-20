package app

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func Test_When_Successful_Should_Return_200_OK(t *testing.T) {
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

func Test_When_NoCustomers_Should_Return_Empty_Array(t *testing.T) {
	// customers := []dto.CustomerResponse{}

	router := mux.NewRouter()
	ch := CustomerHandler{}
	router.HandleFunc("/customers", ch.GetAllCustomers)

	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)

	response := httptest.NewRecorder()
	router.ServeHTTP(response, request)

	if body := response.Body.String(); body != "[]" {
		t.Errorf("Expected an empty array. Got '%s'", body)
	}
}

/*
	Remaining tests:
	0. Return 200 OK on success
	1. Return empty array
	2. Return a single customer
	3. Return multiple customer
	4. Return proper error message when upstream failed
*/
