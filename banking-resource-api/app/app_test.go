package app

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
)

func VerifyIfRouteAvailable(r Route, router *mux.Router, t *testing.T) {
	route := router.GetRoute(r.Name())

	if route == nil {
		t.Fatalf("Route '%v' not available", r.Name())
	}

	path, _ := route.GetPathTemplate()
	if path != r.PathTemplate() {
		t.Fatalf("Route Path is incorrect. Expected: '%v', Received: '%v'", r.PathTemplate(), path)
	}

	handler := route.GetHandler()
	if handler == nil {
		t.Fatalf("Unable to find handler for route: '%v'", r.Name())
	}
}

func Test_Given_DefaultApplication_When_RouterIsSetup_Then_GetAllCustomersRouteIsAvailable(t *testing.T) {
	defaultApp := DefaultApplication{}
	router := defaultApp.SetupRouter()

	VerifyIfRouteAvailable(Route(GetAllCustomers), router, t)
}

func verifyPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("Panic didn't happen")
	}
}

func inducePanicToPreemptOSExit() error {
	return nil
}

func DummyOpenSql(string, string) (*sqlx.DB, error) {
	return nil, nil
}

func Test_Given_DefaultApplication_When_Started_Then_ListenAndServeShouldUseHostPortFromOsEnvVars(t *testing.T) {
	t.Setenv("API_HOST", "localhost")
	t.Setenv("API_PORT", "8080")
	t.Cleanup(func() {
		t.Setenv("API_HOST", "")
		t.Setenv("API_PORT", "")
	})

	defer verifyPanic(t) // Needed to recover from inducePanicToPreemptOSExit

	testApp := DefaultApplication{
		func(addr string, h http.Handler) error {
			if addr != "localhost:8080" {
				t.Errorf("ListenAndServe received: '%v'", addr)
			}
			return inducePanicToPreemptOSExit()
		}, DummyOpenSql,
	}
	Start(testApp)
}

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
func DummyListenAndServe(addr string, h http.Handler) error {
	return nil
}

func Test_Given_DefaultApplicationWithMissingHostEnvVars_When_Started_Then_ItPanicsWithCorrectDetails(t *testing.T) {
	defer verifyPanicWithMessage(t, "Env variable API_HOST not found")

	testApp := NewDefaultApplication(DummyListenAndServe, DummyOpenSql)

	Start(testApp)
}

func Test_Given_DefaultApplicationWithMissingPortEnvVars_When_Started_Then_ItPanicsWithCorrectDetails(t *testing.T) {
	t.Setenv("API_HOST", "localhost")
	t.Cleanup(func() {
		t.Setenv("API_HOST", "")
	})

	defer verifyPanicWithMessage(t, "Env variable API_PORT not found")

	testApp := NewDefaultApplication(DummyListenAndServe, DummyOpenSql)

	Start(testApp)
}
