package domain

type CustomerRepositoryStub struct {
	customers []Customer
}

func (s CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return s.customers, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{"1", "Shankar", "Chennai", "560048", "08-Nov-1983", "Active"},
		{"2", "Shankya", "Trivandrum", "560048", "18-Feb-2010", "Active"},
		{"3", "Dhyan ", "Bangalore", "560048", "30-Sep-2013", "Active"},
	}
	return CustomerRepositoryStub{customers}
}
