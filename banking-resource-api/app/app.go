package app

import (
	"banking-resource-api/logger"
	"banking-resource-api/service"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Application interface {
	SetupRouter() *mux.Router
	ListenAndServeRoutes(*mux.Router, string, string)
}

type DefaultApplication struct {
	ListenAndServe func(string, http.Handler) error
}

func (a DefaultApplication) SetupRouter() *mux.Router {
	ch := CustomerHandler{Service: service.NewCustomerService()}
	const GetAllCustomersRoute = Route(GetAllCustomers)

	router := mux.NewRouter()
	router.HandleFunc(GetAllCustomersRoute.PathTemplate(), ch.GetAllCustomers).
		Name(GetAllCustomersRoute.Name())

	return router
}

func (a DefaultApplication) ListenAndServeRoutes(router *mux.Router, host string, port string) {
	err := a.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router)
	logger.Fatal("Exiting server: " + err.Error())
}

func Start(a Application) {
	fmt.Println("Starting banking resource service")

	// Define routes
	router := a.SetupRouter()

	// Start server
	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")
	a.ListenAndServeRoutes(router, host, port)
}
