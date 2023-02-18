package domain

import "backend/errs"

type Customer struct {
	Id          string `db:"cust_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"dob"`
	Status      string
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	ById(string) (*Customer, *errs.AppError)
}
