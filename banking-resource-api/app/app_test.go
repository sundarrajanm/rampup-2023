package app

import (
	"net/http"
	"testing"

	"github.com/gorilla/mux"
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

func Test_DefaultApplication_Should_Have_GetAllCustomers_Route(t *testing.T) {
	defaultApp := DefaultApplication{}
	router := defaultApp.SetupRouter()

	VerifyIfRouteAvailable(Route(GetAllCustomers), router, t)
}

func DummyListenAndServe(addr string, handler http.Handler) error {
	return nil
}

func verifyPanic(t *testing.T) {
	if r := recover(); r == nil {
		t.Errorf("The code did not panic")
	}
}

func Test_DefaultApplication_Should_ListenAndServe_FromHostPort_Setup_InEnv(t *testing.T) {
	t.Setenv("API_HOST", "localhost")
	t.Setenv("API_PORT", "8080")

	defer verifyPanic(t)

	testApp := DefaultApplication{
		func(addr string, h http.Handler) error {
			if addr != "localhost:8080" {
				t.Fatalf("ListenAndServe received: '%v'", addr)
			}
			return nil
		},
	}
	Start(testApp)
}

// func Test_DefaultApplication_Should_Fail_In_ListenAndServe_WhenHostPort_UnAvailable_InEnv(t *testing.T) {

// }
