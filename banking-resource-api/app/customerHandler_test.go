package app

import (
	"banking-resource-api/dto"
	"banking-resource-api/errs"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

// //////////// CustomerService Mock ////////////
type DummyCustomerService struct {
	getAllCustomersMock func() ([]dto.CustomerResponse, *errs.AppError)
}

func (d DummyCustomerService) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	return d.getAllCustomersMock()
}

// //////////////////////////////////////////////

func executeWithMockCustomerServiceResponse(mock func() ([]dto.CustomerResponse, *errs.AppError)) *httptest.ResponseRecorder {
	// Arrange
	router := mux.NewRouter()
	ch := CustomerHandler{DummyCustomerService{mock}}
	router.HandleFunc("/customers", ch.GetAllCustomers)
	request, _ := http.NewRequest(http.MethodGet, "/customers", nil)
	responseWriter := httptest.NewRecorder()

	// Act
	router.ServeHTTP(responseWriter, request)
	return responseWriter
}

func Test_When_Successful_Should_Return_200_OK(t *testing.T) {
	// Arrange
	getAllCustomersMock := func() ([]dto.CustomerResponse, *errs.AppError) {
		return []dto.CustomerResponse{}, nil
	}

	// Act
	response := executeWithMockCustomerServiceResponse(getAllCustomersMock)

	// Assert
	if response.Code != http.StatusOK {
		t.Errorf("Got response code: %d", response.Code)
	}
}

func Test_When_NoCustomers_Should_Return_Empty_Array(t *testing.T) {
	// Arrange
	var result []dto.CustomerResponse
	getAllCustomersMock := func() ([]dto.CustomerResponse, *errs.AppError) {
		return result, nil
	}

	// Act
	response := executeWithMockCustomerServiceResponse(getAllCustomersMock)

	// Assert
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
	getAllCustomersMock := func() ([]dto.CustomerResponse, *errs.AppError) {
		return []dto.CustomerResponse{{}, {}, {}}, nil
	}

	// Act
	response := executeWithMockCustomerServiceResponse(getAllCustomersMock)

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

func Test_When_Service_Failed_Should_Return_AppError(t *testing.T) {
	// Arrange
	getAllCustomersMock := func() ([]dto.CustomerResponse, *errs.AppError) {
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// Act
	response := executeWithMockCustomerServiceResponse(getAllCustomersMock)

	// Assert
	var errResponse errs.AppError
	err := json.NewDecoder(response.Body).Decode(&errResponse)

	if err != nil {
		t.Fatalf("Unable to parse error response from server %q into AppError, '%v'", response.Body, err)
	}

	if errResponse.Code != http.StatusInternalServerError {
		t.Fatalf("Service failure is an internal server error but received: '%v'", errResponse.Code)
	}

	if errResponse.Message != "Unexpected database error" {
		t.Fatalf("Service failure message is not correct: '%v'", errResponse.Message)
	}
}

/*
	Remaining tests:
	0. Return 200 OK on success - DONE
	1. Return empty array - DONE
	2. Return customer array on success - DONE
	3. Return proper error message when upstream failed
*/
