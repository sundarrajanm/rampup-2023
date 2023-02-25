package service

import (
	"banking-resource-api/domain"
	"banking-resource-api/dto"
	"banking-resource-api/errs"
	"banking-resource-api/logger"
)

type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError)
	GetCustomerById(string) (dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (d DefaultCustomerService) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	logger.Info("Enter DefaultCustomerService: GetAllCustomers")
	customers, err := d.repo.FindAll()
	if err != nil {
		return nil, err
	}

	customersDTO := make([]dto.CustomerResponse, 0)
	for _, c := range customers {
		customersDTO = append(customersDTO, *c.ToDTO())
	}
	logger.Info("Exiting DefaultCustomerService: GetAllCustomers")
	return customersDTO, nil
}

func (d DefaultCustomerService) GetCustomerById(string) (dto.CustomerResponse, *errs.AppError) {
	return dto.CustomerResponse{}, nil
}

func NewCustomerService(repo domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repo}
}
