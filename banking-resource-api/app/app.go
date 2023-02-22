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
	logger.Fatal(err.Error())
}

func checkMandatoryEnvVars(vars ...string) {
	for _, v := range vars {
		value := os.Getenv(v)
		logger.Info("Checking: " + v + ", value: " + value)
		if value == "" {
			errMsg := "Env variable " + v + " not found"
			logger.Error(errMsg)
			panic(errMsg)
		}
	}
}

func Start(a Application) {
	logger.Info("Starting banking resource service")
	checkMandatoryEnvVars("API_HOST", "API_PORT")

	// Define routes
	router := a.SetupRouter()

	// Start server
	host := os.Getenv("API_HOST")
	port := os.Getenv("API_PORT")
	a.ListenAndServeRoutes(router, host, port)
}
