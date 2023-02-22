package app

type Route int

const (
	GetAllCustomers = iota
)

func (r Route) Name() string {
	return [...]string{
		"GetAllCustomers",
	}[r]
}

func (r Route) PathTemplate() string {
	return [...]string{
		"/customers",
	}[r]
}
