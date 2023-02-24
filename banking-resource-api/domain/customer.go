package domain

import "banking-resource-api/dto"

type Customer struct {
	Id          string `db:"cust_id"`
	Name        string
	City        string
	Zipcode     string
	DateofBirth string `db:"dob"`
	Status      string
}

func (c Customer) statusAsText() string {
	statusAsText := "active"
	if c.Status == "0" {
		statusAsText = "inactive"
	}
	return statusAsText
}

func (c Customer) ToDTO() *dto.CustomerResponse {

	return &dto.CustomerResponse{
		Id:          c.Id,
		Name:        c.Name,
		City:        c.City,
		Zipcode:     c.Zipcode,
		DateofBirth: c.DateofBirth,
		Status:      c.statusAsText(),
	}
}
