package service

import (
	"banking-resource-api/domain"
	"banking-resource-api/dto"
	"banking-resource-api/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError)
	GetCustomerById(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := d.repo.FindAll()
	if err != nil {
		return nil, err
	}

	customersDTO := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		customersDTO = append(customersDTO, *c.ToDTO())
	}
	return customersDTO, nil
}

func (d DefaultCustomerService) GetCustomerById(id string) (*dto.CustomerResponse, *errs.AppError) {
	customer, appError := d.repo.FindById(id)

	if appError != nil {
		return nil, appError
	}

	return customer.ToDTO(), nil
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
