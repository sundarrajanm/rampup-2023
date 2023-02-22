package app

import (
	"banking-resource-api/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Application interface {
	SetupRouter() *mux.Router
	ListenAndServeRoutes(*mux.Router, string, string)
}

type DefaultApplication struct{}

func (a DefaultApplication) SetupRouter() *mux.Router {
	ch := CustomerHandler{Service: service.NewCustomerService()}
	const GetAllCustomersRoute = Route(GetAllCustomers)

	router := mux.NewRouter()
	router.HandleFunc(GetAllCustomersRoute.PathTemplate(), ch.GetAllCustomers).
		Name(GetAllCustomersRoute.Name())

	return router
}

func (a DefaultApplication) ListenAndServeRoutes(router *mux.Router, host string, port string) {
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router))
}

func Start(a Application) {
	fmt.Println("Starting banking resource service")

	// Define routes
	router := a.SetupRouter()

	// Start server
	host := "localhost"
	port := "8000"
	a.ListenAndServeRoutes(router, host, port)
}
