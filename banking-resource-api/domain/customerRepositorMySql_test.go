package domain

import (
	"testing"

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

func Test_Given_FindAll_Then_UseMySqlDriver(t *testing.T) {
	setAllEnvVars(t)

	NewCustomerRepoMySql(func(driver string, connectionString string) (*sqlx.DB, error) {
		if driver != "mysql" {
			t.Errorf("Unknown driver used: '%v'", "something")
		}
		return nil, nil
	})
}
