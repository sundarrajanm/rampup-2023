package domain

import "banking-resource-api/errs"

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
}

type DefaultCustomerRepo struct{}

func (d DefaultCustomerRepo) FindAll() ([]Customer, *errs.AppError) {
	return []Customer{}, nil
}

func NewCustomerRepository() CustomerRepository {
	return DefaultCustomerRepo{}
}
