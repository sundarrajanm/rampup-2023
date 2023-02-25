package controller

type Route int

const (
	GetAllCustomers = iota
	GetCustomerById
)

func (r Route) Name() string {
	return [...]string{
		"GetAllCustomers",
		"GetCustomerById",
	}[r]
}

func (r Route) PathTemplate() string {
	return [...]string{
		"/customers",
		"/customers/{customer_id:[0-9]+}",
	}[r]
}
