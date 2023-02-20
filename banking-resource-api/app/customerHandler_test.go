package app

import (
	"encoding/json"
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

	var result []any
	err := json.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of any, '%v'", response.Body, err)
	}

	if len(result) != 0 {
		t.Fatalf("Result was not empty, found: '%v'", result)
	}
}

/*
	Remaining tests:
	0. Return 200 OK on success - DONE
	1. Return empty array - DONE
	2. Return a single customer
	3. Return multiple customer
	4. Return proper error message when upstream failed
*/
