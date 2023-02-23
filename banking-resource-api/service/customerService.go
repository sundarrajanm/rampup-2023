package service

import (
	"banking-resource-api/domain"
	"banking-resource-api/dto"
	"banking-resource-api/errs"
	"banking-resource-api/logger"
)

type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	logger.Info("Inside customer server - get all customers")
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

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
