package service

import (
	"backend/domain"
	"backend/dto"
	"backend/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	customers, err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	customersDTO := make([]dto.CustomerResponse, 0)

	for _, c := range customers {
		customersDTO = append(customersDTO, *c.ToDTO())
	}

	return customersDTO, nil
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	customer, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	return customer.ToDTO(), nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
