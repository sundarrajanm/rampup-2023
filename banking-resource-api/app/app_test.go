package app

import (
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

func Test_DefaultApplication_Should_ListenAndServe_FromHostPort_Setup_InEnv() {
	defaultApp := DefaultApplication{}
}
