package app

import (
	"banking-resource-api/dto"
	"banking-resource-api/errs"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func (d DummyCustomerService) GetCustomerById(id string) (dto.CustomerResponse, *errs.AppError) {
	return d.getCustomerByIdMock(id)
}

func getCustomerByIdMockShouldReturn(response dto.CustomerResponse, err *errs.AppError) func(string) (dto.CustomerResponse, *errs.AppError) {
	return func(string) (dto.CustomerResponse, *errs.AppError) {
		return response, err
	}
}

func executeWithMockGetCustomerById(mock func(string) (dto.CustomerResponse, *errs.AppError)) *httptest.ResponseRecorder {
	// Arrange
	router := mux.NewRouter()
	ch := CustomerHandler{DummyCustomerService{getCustomerByIdMock: mock}}

	route := Route(GetCustomerById)
	router.HandleFunc(route.PathTemplate(), ch.GetCustomerById)
	request, _ := http.NewRequest(http.MethodGet, "/customers/1000", nil)
	responseWriter := httptest.NewRecorder()

	// Act
	router.ServeHTTP(responseWriter, request)
	return responseWriter
}

func Test_Given_GetCustomerByIdRequest_When_Successful_Then_Return200OK_(t *testing.T) {
	// Arrange
	mock := getCustomerByIdMockShouldReturn(dto.CustomerResponse{}, nil)

	// Act
	response := executeWithMockGetCustomerById(mock)

	// Assert
	if response.Code != http.StatusOK {
		t.Errorf("Got response code: %d", response.Code)
	}

	contentTypeHeader := response.Result().Header.Get("Content-Type")
	if contentTypeHeader != "application/json" {
		t.Errorf("Content-Type was not application/json, received: %s", contentTypeHeader)
	}
}
