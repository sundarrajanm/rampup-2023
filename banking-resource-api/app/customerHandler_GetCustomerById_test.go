package app

import (
	"banking-resource-api/dto"
	"banking-resource-api/errs"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func (d DummyCustomerService) GetCustomerById(id string) (*dto.CustomerResponse, *errs.AppError) {
	return d.getCustomerByIdMock(id)
}

func getCustomerByIdMockShouldReturn(response *dto.CustomerResponse, err *errs.AppError) func(string) (*dto.CustomerResponse, *errs.AppError) {
	return func(string) (*dto.CustomerResponse, *errs.AppError) {
		return response, err
	}
}

const CustomerId1000 = "1000"

func executeWithMockGetCustomerById(mock func(string) (*dto.CustomerResponse, *errs.AppError)) *httptest.ResponseRecorder {
	// Arrange
	router := mux.NewRouter()
	ch := CustomerHandler{DummyCustomerService{getCustomerByIdMock: mock}}

	route := Route(GetCustomerById)
	router.HandleFunc(route.PathTemplate(), ch.GetCustomerById)
	request, _ := http.NewRequest(http.MethodGet, "/customers/"+CustomerId1000, nil)
	responseWriter := httptest.NewRecorder()

	// Act
	router.ServeHTTP(responseWriter, request)
	return responseWriter
}

func Test_Given_GetCustomerByIdRequest_When_Successful_Then_ReturnCustomerWith200OK(t *testing.T) {
	expectedCustomerDTO := &dto.CustomerResponse{
		Id:          "10",
		Name:        "Bob",
		City:        "Bangalore",
		Zipcode:     "560048",
		DateofBirth: "10-10-1978",
		Status:      "active",
	}

	// Arrange
	mock := getCustomerByIdMockShouldReturn(expectedCustomerDTO, nil)

	// Act
	response := executeWithMockGetCustomerById(mock)

	// Assert
	if response.Code != http.StatusOK {
		t.Errorf("expected: 200, receieved: %d", response.Code)
	}

	contentTypeHeader := response.Result().Header.Get("Content-Type")
	if contentTypeHeader != "application/json" {
		t.Errorf("Content-Type was not application/json, received: %s", contentTypeHeader)
	}

	var result dto.CustomerResponse
	err := json.NewDecoder(response.Body).Decode(&result)
	if err != nil {
		t.Errorf("Response Body JSON parsing failed")
	}

	if result != *expectedCustomerDTO {
		t.Errorf(fmt.Sprintf("Expected Body: '%v', Received: '%v'", expectedCustomerDTO,
			result))
	}
}

func Test_Given_GetCustomerByIdRequest_When_NotFound_Then_Return404(t *testing.T) {
	expectedErrorResponse := errs.NewNotFoundError("Customer with Id " + CustomerId1000 + " not found")
	validateErrorResponse(*expectedErrorResponse, t)
}

func Test_Given_GetCustomerByIdRequest_When_ServiceInternallyFailed_Then_Return500(t *testing.T) {
	expectedErrorResponse := errs.NewUnexpectedError("Unexpected database error")
	validateErrorResponse(*expectedErrorResponse, t)
}

func validateErrorResponse(expectedError errs.AppError, t *testing.T) {
	// Arrange
	mock := getCustomerByIdMockShouldReturn(nil, &expectedError)

	// Act
	response := executeWithMockGetCustomerById(mock)

	// Assert
	if response.Code != expectedError.Code {
		t.Errorf("expected: '%d', received: '%d'", expectedError.Code, response.Code)
	}

	contentTypeHeader := response.Result().Header.Get("Content-Type")
	if contentTypeHeader != "application/json" {
		t.Errorf("Content-Type was not application/json, received: %s", contentTypeHeader)
	}

	var errResponse errs.AppError
	err := json.NewDecoder(response.Body).Decode(&errResponse)

	if err != nil {
		t.Fatalf("Unable to parse error response from server %q into AppError, '%v'", response.Body, err)
	}

	if errResponse != expectedError {
		t.Errorf(fmt.Sprintf("Expected Body: '%v', Received: '%v'", expectedError, errResponse))
	}
}
