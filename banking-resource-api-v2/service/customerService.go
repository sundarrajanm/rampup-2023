package service

import "banking-resource-api-v2/dto"

//go:generate mockgen -destination=../mocks/service/mockCustomerService.go -package=service banking-resource-api-v2 CustomerService
type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, error)
}
