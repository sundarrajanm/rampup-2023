package service

import (
	"banking-resource-api/domain"
	"banking-resource-api/dto"
	"banking-resource-api/errs"
	"testing"
)

func (d DummyTestRepo) FindById(id string) (*domain.Customer, *errs.AppError) {
	return d.getCustomerByIdMock(id)
}

func Test_GivenGetCustomersById_WhenSuccessful_ThenReturn_CustomerResponseDTO(t *testing.T) {
	expectedCustomerDTO := dto.CustomerResponse{
		Id:          "1",
		Name:        "Test",
		City:        "Bengaluru",
		Zipcode:     "560048",
		DateofBirth: "01-01-1947",
		Status:      "active",
	}

	service := NewCustomerService(DummyTestRepo{
		getCustomerByIdMock: func(string) (*domain.Customer, *errs.AppError) {
			return &domain.Customer{
				Id:          "1",
				Name:        "Test",
				City:        "Bengaluru",
				Zipcode:     "560048",
				DateofBirth: "01-01-1947",
				Status:      "1",
			}, nil
		},
	})

	customer, _ := service.GetCustomerById("1000")

	if customer == nil {
		t.Fatalf("Expected: customer, Received: nil")
	}

	if *customer != expectedCustomerDTO {
		t.Errorf("Expected: '%v', Received: '%v'", expectedCustomerDTO, *customer)
	}
}

func Test_GivenGetCustomerById_WhenSuccessful_ThenForStatus_0_MapTo_Inactive(t *testing.T) {
	service := NewCustomerService(DummyTestRepo{
		getAllCustomersMock: func() ([]domain.Customer, *errs.AppError) {
			return []domain.Customer{{
				Id:          "1",
				Name:        "Test",
				City:        "Bengaluru",
				Zipcode:     "560048",
				DateofBirth: "01-01-1947",
				Status:      "0",
			}}, nil
		},
	})
	customers, _ := service.GetAllCustomers()

	if customers[0].Status != "inactive" {
		t.Errorf("Expected: inactive, Received: '%v'", customers[0].Status)
	}
}

func verifyExpectedError(expectedError *errs.AppError, t *testing.T) {
	service := NewCustomerService(DummyTestRepo{
		getCustomerByIdMock: func(string) (*domain.Customer, *errs.AppError) {
			return nil, expectedError
		},
	})

	_, err := service.GetCustomerById("1000")

	if err == nil {
		t.Errorf("Expected appError but no err found")
		return
	}

	if err != expectedError {
		t.Errorf("Expected: '%v', Received: '%v'", expectedError, err)
		return
	}
}

func Test_GivenGetCustomerById_WhenNoCustomerFound_ThenReturn_CorrectAppError(t *testing.T) {
	verifyExpectedError(errs.NewNotFoundError("Customer with id: 1000 not found"), t)
}

func Test_GivenGetCustomerById_WhenRepoQueryFetchFailed_ThenReturn_CorrectAppError(t *testing.T) {
	verifyExpectedError(errs.NewUnexpectedError("Unexpected database error"), t)
}
