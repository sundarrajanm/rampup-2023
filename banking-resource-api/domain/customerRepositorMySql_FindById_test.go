package domain

import (
	"banking-resource-api/errs"
	"database/sql"
	"errors"
	"fmt"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

const CustomerId string = "1"

func Test_Given_FindById_WhenCustomerFound_ThenReturnDomainCustomer(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"cust_id", "name", "city", "zipcode", "dob", "status"}).
		AddRow(CustomerId, "Shankya", "Bengaluru", "560048", "18-02-2010", "1")
	mock.ExpectQuery("select cust_id, name, city, zipcode, dob, status from customers where cust_id = ?").
		WithArgs(CustomerId).
		WillReturnRows(rows)

	setAllEnvVars(t)
	repo := NewCustomerRepoMySql(func(s1, s2 string) (*sqlx.DB, error) { return sqlx.NewDb(db, "mysql"), nil })

	customer, _ := repo.FindById(CustomerId)
	expectedCustomer := &Customer{
		Id:          CustomerId,
		Name:        "Shankya",
		City:        "Bengaluru",
		Zipcode:     "560048",
		DateofBirth: "18-02-2010",
		Status:      "1",
	}

	if *expectedCustomer != *customer {
		t.Errorf("Expected: '%v', Received: '%v'", expectedCustomer, customer)
	}
}

func verifyErrorScenario(dbError error, expectedError *errs.AppError, t *testing.T) {
	setAllEnvVars(t)
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectQuery("select cust_id, name, city, zipcode, dob, status from customers where cust_id = ?").
		WithArgs(CustomerId).
		WillReturnError(dbError)

	repo := NewCustomerRepoMySql(func(s1, s2 string) (*sqlx.DB, error) { return sqlx.NewDb(db, "mysql"), nil })

	_, appError := repo.FindById(CustomerId)
	if *expectedError != *appError {
		t.Errorf("Expected: '%v', Received: '%v'", expectedError, appError)
	}
}

func Test_Given_FindById_WhenFailed_ThenReturnCorrectAppError(t *testing.T) {
	verifyErrorScenario(sql.ErrNoRows, errs.NewNotFoundError(fmt.Sprintf("Customer with Id: '%s' not found", CustomerId)), t)
	verifyErrorScenario(errors.New("Network connection failed"), errs.NewUnexpectedError("Unexpected database error"), t)
}
