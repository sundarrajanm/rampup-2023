package domain

import "banking-resource-api/errs"

// Simple In Memory DB to avoid SQL setup
type CustomerRepoInMemory struct{}

func (d CustomerRepoInMemory) FindAll() ([]Customer, *errs.AppError) {
	return []Customer{}, nil
}

func NewCustomerRepositoryInMemory() CustomerRepository {
	return CustomerRepoInMemory{}
}
