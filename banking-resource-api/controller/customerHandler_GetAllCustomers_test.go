package controller

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
	getCustomerByIdMock func(string) (*dto.CustomerResponse, *errs.AppError)
}

func (d DummyCustomerService) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	return d.getAllCustomersMock()
}

// //////////////////////////////////////////////

func getAllCustomersMockShouldReturn(response []dto.CustomerResponse, err *errs.AppError) func() ([]dto.CustomerResponse, *errs.AppError) {
	return func() ([]dto.CustomerResponse, *errs.AppError) {
		return response, err
	}
}

func executeWithMockCustomerServiceResponse(mock func() ([]dto.CustomerResponse, *errs.AppError)) *httptest.ResponseRecorder {
	// Arrange
	router := mux.NewRouter()
	ch := CustomerHandler{DummyCustomerService{getAllCustomersMock: mock}}

	route := Route(GetAllCustomers)
	router.HandleFunc(route.PathTemplate(), ch.GetAllCustomers)
	request, _ := http.NewRequest(http.MethodGet, route.PathTemplate(), nil)
	responseWriter := httptest.NewRecorder()

	// Act
	router.ServeHTTP(responseWriter, request)
	return responseWriter
}

func Test_Given_GetAllCustomersRequest_When_Successful_Then_Return200OK_(t *testing.T) {
	// Arrange
	mock := getAllCustomersMockShouldReturn([]dto.CustomerResponse{}, nil)

	// Act
	response := executeWithMockCustomerServiceResponse(mock)

	// Assert
	if response.Code != http.StatusOK {
		t.Errorf("Got response code: %d", response.Code)
	}
}

func Test_Given_GetAllCustomersRequest_When_NoCustomers_Then_ReturnEmptyArray(t *testing.T) {
	// Arrange
	var result []dto.CustomerResponse
	mock := getAllCustomersMockShouldReturn(result, nil)

	// Act
	response := executeWithMockCustomerServiceResponse(mock)

	// Assert
	err := json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of any, '%v'", response.Body, err)
	}
	if len(result) != 0 {
		t.Fatalf("Result was not empty, found: '%v'", result)
	}
}

func Test_Given_GetAllCustomersRequest_WhenThereAreCustomers_Then_ReturnArrayOfCustomers(t *testing.T) {
	// Arrange
	mock := getAllCustomersMockShouldReturn([]dto.CustomerResponse{{}, {}, {}}, nil)

	// Act
	response := executeWithMockCustomerServiceResponse(mock)

	// Assert
	var result []dto.CustomerResponse
	err := json.NewDecoder(response.Body).Decode(&result)

	if err != nil {
		t.Errorf("Unable to parse response from server %q into slice of any, '%v'", response.Body, err)
	}

	if len(result) != 3 {
		t.Errorf("Expected 3 customers, found: '%v'", result)
	}

	expectedContentType := "application/json"
	actualContentType := response.Result().Header.Get("Content-Type")
	if actualContentType != expectedContentType {
		t.Errorf("Expected: '%v', Received: '%v'", expectedContentType, actualContentType)
	}
}

func Test_Given_GetAllCustomersRequest_When_ServiceInternallyFailed_Then_ReturnAppError(t *testing.T) {
	// Arrange
	mock := getAllCustomersMockShouldReturn(nil, errs.NewUnexpectedError("Unexpected database error"))

	// Act
	response := executeWithMockCustomerServiceResponse(mock)

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
