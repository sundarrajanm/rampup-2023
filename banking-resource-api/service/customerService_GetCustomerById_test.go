package service

import (
	"banking-resource-api/domain"
	"banking-resource-api/errs"
	"testing"
)

func (d DummyTestRepo) FindById(id string) (*domain.Customer, *errs.AppError) {
	return d.getCustomerByIdMock(id)
}

func Test_GivenGetCustomersById_WhenSuccessful_ThenReturn_CustomerResponseDTO(t *testing.T) {
	service := NewCustomerService(DummyTestRepo{
		getCustomerByIdMock: func(string) (*domain.Customer, *errs.AppError) {
			return &domain.Customer{
				Id:          "1",
				Name:        "Test",
				City:        "Bengaluru",
				Zipcode:     "560048",
				DateofBirth: "01-01-1947",
				Status:      "active",
			}, nil
		},
	})
	customers, _ := service.GetCustomerById("1000")

	if customers == nil {
		t.Errorf("Expected: customer, Received: nil")
	}
}

// func Test_GivenGetAllCustomers_WhenSuccessful_ThenForStatus_0_MapTo_Inactive(t *testing.T) {
// 	service := NewCustomerService(DummyTestRepo{
// 		getAllCustomersMock: func() ([]domain.Customer, *errs.AppError) {
// 			return []domain.Customer{{
// 				Id:          "1",
// 				Name:        "Test",
// 				City:        "Bengaluru",
// 				Zipcode:     "560048",
// 				DateofBirth: "01-01-1947",
// 				Status:      "0",
// 			}}, nil
// 		},
// 	})
// 	customers, _ := service.GetAllCustomers()

// 	if customers[0].Status != "inactive" {
// 		t.Errorf("Expected: inactive, Received: '%v'", customers[0].Status)
// 	}
// }

// func Test_GivenGetAllCustomers_WhenSuccessful_ThenReturn_EmptyCustomerResponseDTO(t *testing.T) {
// 	service := NewCustomerService(DummyTestRepo{
// 		getAllCustomersMock: func() ([]domain.Customer, *errs.AppError) {
// 			return []domain.Customer{}, nil
// 		},
// 	})
// 	customers, _ := service.GetAllCustomers()

// 	if len(customers) != 0 {
// 		t.Errorf("Expected: 0, Received: '%v'", len(customers))
// 	}
// }

// func Test_GivenGetAllCustomers_WhenFailed_ThenReturn_AppError(t *testing.T) {
// 	service := NewCustomerService(DummyTestRepo{
// 		getAllCustomersMock: func() ([]domain.Customer, *errs.AppError) {
// 			return nil, errs.NewUnexpectedError("Unexpected database error")
// 		},
// 	})

// 	_, err := service.GetAllCustomers()

// 	if err == nil {
// 		t.Errorf("Expected appError")
// 		return
// 	}

// 	if err.Code != http.StatusInternalServerError {
// 		t.Errorf("Expected: '%v', Received: '%v'", http.StatusInternalServerError, err.Code)
// 		return
// 	}

// 	if err.Message != "Unexpected database error" {
// 		t.Errorf("Expected: 'Unexpected database error', Received: '%v'", err.Message)
// 	}
// }
