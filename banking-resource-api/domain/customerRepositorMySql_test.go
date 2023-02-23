package domain

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
)

func verifyPanicWithMessage(t *testing.T, msg string) {
	r := recover()

	t.Logf("Panic message: '%v'", r)

	if r == nil {
		t.Errorf("Panic didn't happen")
	}

	if r != msg {
		t.Errorf("Expected: '%v', Received: '%v'", msg, r)
	}
}
func setAllEnvVars(t *testing.T) {
	t.Setenv("DB_USER", "root")
	t.Setenv("DB_PASSWORD", "secret")
	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_PORT", "3306")
	t.Setenv("DB_NAME", "banking")
	t.Cleanup(func() {
		t.Setenv("DB_USER", "")
		t.Setenv("DB_PASSWORD", "")
		t.Setenv("DB_HOST", "")
		t.Setenv("DB_PORT", "")
		t.Setenv("DB_NAME", "")
	})
}

func makeEnvVarEmpty(varName string, t *testing.T) {
	setAllEnvVars(t)
	t.Setenv(varName, "")
}

func Test_Given_DB_USER_EnvVar_WhenEmpty_ItPanicsWithCorrectDetails(t *testing.T) {
	defer verifyPanicWithMessage(t, "Env variable DB_USER not found")
	makeEnvVarEmpty("DB_USER", t)
	NewCustomerRepoMySql(nil)
}

func Test_Given_DB_PASSWORD_EnvVar_WhenEmpty_ItPanicsWithCorrectDetails(t *testing.T) {
	defer verifyPanicWithMessage(t, "Env variable DB_PASSWORD not found")
	makeEnvVarEmpty("DB_PASSWORD", t)
	NewCustomerRepoMySql(nil)
}

func Test_Given_DB_HOST_EnvVar_WhenEmpty_ItPanicsWithCorrectDetails(t *testing.T) {
	defer verifyPanicWithMessage(t, "Env variable DB_HOST not found")
	makeEnvVarEmpty("DB_HOST", t)
	NewCustomerRepoMySql(nil)
}

func Test_Given_DB_PORT_EnvVar_WhenEmpty_ItPanicsWithCorrectDetails(t *testing.T) {
	defer verifyPanicWithMessage(t, "Env variable DB_PORT not found")
	makeEnvVarEmpty("DB_PORT", t)
	NewCustomerRepoMySql(nil)
}

func Test_Given_DB_NAME_EnvVar_WhenEmpty_ItPanicsWithCorrectDetails(t *testing.T) {
	defer verifyPanicWithMessage(t, "Env variable DB_NAME not found")
	makeEnvVarEmpty("DB_NAME", t)
	NewCustomerRepoMySql(nil)
}

func Test_Given_FindAll_Then_UseMySqlDriver_And_CorrectConnectionString(t *testing.T) {
	setAllEnvVars(t)
	expectedConnectionString := GetConnectionString()

	NewCustomerRepoMySql(func(driver string, connectionString string) (*sqlx.DB, error) {
		if driver != MySqlDriver {
			t.Errorf("Expected db driver: '%v', Received: '%v'", MySqlDriver, driver)
		}
		if connectionString != expectedConnectionString {
			t.Errorf("Expected connection string: '%v', Received: '%v'", expectedConnectionString,
				connectionString)
		}
		return nil, nil
	})
}

func Test_Given_FindAll_WhenCustomersInDB_ThenReturnDomainCustomers(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	rows := sqlmock.NewRows([]string{"cust_id", "name", "city", "zipcode", "dob", "status"}).
		AddRow("1", "Shankya", "Bengaluru", "560048", "18-02-2010", "1")
	mock.ExpectQuery("select cust_id, name, city, zipcode, dob, status from customers").
		WillReturnRows(rows)

	setAllEnvVars(t)
	repo := NewCustomerRepoMySql(func(s1, s2 string) (*sqlx.DB, error) { return sqlx.NewDb(db, "mysql"), nil })

	customers, _ := repo.FindAll()
	if len(customers) != 1 {
		t.Errorf("Expected: 1, Received: '%d'", len(customers))
	}
}
