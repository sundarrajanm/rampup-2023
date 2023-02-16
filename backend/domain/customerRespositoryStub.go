package domain

import (
	"backend/errs"
	"log"
)

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func (s CustomerRepositoryStub) ById(id string) (*Customer, *errs.AppError) {
	if id == "1" {
		return &Customer{
			"1", "Shankar", "Chennai", "560048", "08-Nov-1983", "Active",
		}, nil
	} else if id == "2" {
		return nil, errs.NewNotFoundError("record not found")
	} else {
		errMsg := "connection to upstream failed"
		log.Println("Error while communicating with DB: " + errMsg)
		return nil, errs.NewUnexpectedError(errMsg)
	}
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "Shankar", "Chennai", "560048", "08-Nov-1983", "Active"},
		{"2", "Shankya", "Trivandrum", "560048", "18-Feb-2010", "Active"},
		{"3", "Dhyan ", "Bangalore", "560048", "30-Sep-2013", "Active"},
	}
	return CustomerRepositoryStub{customers}
}
