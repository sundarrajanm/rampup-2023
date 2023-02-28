package service

import "banking-resource-api-v2/dto"

//go:generate mockgen -destination=../mocks/service/mockCustomerService.go -package=service . CustomerService
type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, error)
}
