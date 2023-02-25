package dto

// swagger:model Customer
type CustomerResponse struct {
	// Id of the Customer
	// in: string
	Id string `json:"customer_id"`

	// Name of the Customer
	// in: string
	Name string `json:"name"`

	// City of the Customer
	// in: string
	City string `json:"city"`

	// Zip Code of the Customer
	// in: string
	Zipcode string `json:"zip_code"`

	// Date of birth of the Customer
	// in: string
	DateofBirth string `json:"date_of_birth"`

	// Status of the customer
	// in: string
	Status string `json:"status"`
}
