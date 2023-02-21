package app

import (
	"banking-resource-api/dto"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// //////////// CustomerService Mock ////////////

var getAllCustomersMock func() []dto.CustomerResponse

type DummyCustomerService struct{}

func (d DummyCustomerService) GetAllCustomers() []dto.CustomerResponse {
	return getAllCustomersMock()
}

// //////////////////////////////////////////////

func Test_When_Successful_Should_Return_200_OK(t *testing.T) {
	// Arrange
	getAllCustomersMock = func() []dto.CustomerResponse {
		return []dto.CustomerResponse{}
	}
	router := mux.NewRouter()
	ch := CustomerHandler{DummyCustomerService{}}
	router.HandleFunc("/customers", ch.GetAllCustomers)
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
	responseWriter := httptest.NewRecorder()

	// Act
	router.ServeHTTP(responseWriter, request)

	// Assert
	if responseWriter.Code != http.StatusOK {
		t.Errorf("Got response code: %d", responseWriter.Code)
	}
}

func Test_When_NoCustomers_Should_Return_Empty_Array(t *testing.T) {
	// Arrange
	getAllCustomersMock = func() []dto.CustomerResponse {
		return []dto.CustomerResponse{}
	}
	ch := CustomerHandler{DummyCustomerService{}}
	router := mux.NewRouter()
	router.HandleFunc("/customers", ch.GetAllCustomers)
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
	response := httptest.NewRecorder()

	// Act
	router.ServeHTTP(response, request)

	// Assert
	var result []dto.CustomerResponse
	err := json.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of any, '%v'", response.Body, err)
	}

	if len(result) != 0 {
		t.Fatalf("Result was not empty, found: '%v'", result)
	}
}

func Test_When_Should_Return_Array_With_Customers(t *testing.T) {
	// Arrange
	getAllCustomersMock = func() []dto.CustomerResponse {
		return []dto.CustomerResponse{{}, {}, {}}
	}
	router := mux.NewRouter()
	ch := CustomerHandler{DummyCustomerService{}}
	router.HandleFunc("/customers", ch.GetAllCustomers)
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
	response := httptest.NewRecorder()

	// Act
	router.ServeHTTP(response, request)

	// Assert
	var result []dto.CustomerResponse
	err := json.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of any, '%v'", response.Body, err)
	}

	if len(result) != 3 {
		t.Fatalf("Expected 3 customers, found: '%v'", result)
	}
}

/*
	Remaining tests:
	0. Return 200 OK on success - DONE
	1. Return empty array - DONE
	2. Return customer array on success - DONE
	3. Return proper error message when upstream failed
*/
