package service

import (
	"banking-resource-api/dto"
	"banking-resource-api/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct{}

func (d DefaultCustomerService) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	return []dto.CustomerResponse{}, nil
}
