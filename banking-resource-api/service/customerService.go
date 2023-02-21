package service

import "banking-resource-api/dto"

type CustomerService interface {
	GetAllCustomers() []dto.CustomerResponse
}

type DefaultCustomerService struct{}

func (d DefaultCustomerService) GetAllCustomers() []dto.CustomerResponse {
	return []dto.CustomerResponse{}
}
