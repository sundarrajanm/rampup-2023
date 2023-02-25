package domain

import "banking-resource-api/errs"

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindById(string) (*Customer, *errs.AppError)
}
