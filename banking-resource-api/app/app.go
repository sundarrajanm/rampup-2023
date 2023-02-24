package app

import (
	"banking-resource-api/logger"
	"banking-resource-api/types"
	"banking-resource-api/utils"
	"fmt"
	"os"

	"github.com/gorilla/mux"
)

type Application interface {
	SetupRouter() *mux.Router
	ListenAndServeRoutes(*mux.Router, string, string)
}

type DefaultApplication struct {
	ListenAndServe  types.HttpListenAndServe
	OpenSql         types.OpenSqlxDB
	CustomerHandler CustomerHandler
}

func (a DefaultApplication) SetupRouter() *mux.Router {
	const GetAllCustomersRoute = Route(GetAllCustomers)
	const GetCustomerByIdRoute = Route(GetCustomerById)

	router := mux.NewRouter()

	router.HandleFunc(GetAllCustomersRoute.PathTemplate(),
		a.CustomerHandler.GetAllCustomers).Name(GetAllCustomersRoute.Name())
	router.HandleFunc(GetCustomerByIdRoute.PathTemplate(),
		a.CustomerHandler.GetCustomerById).Name(GetCustomerByIdRoute.Name())

	return router
}

func (a DefaultApplication) ListenAndServeRoutes(router *mux.Router, host string, port string) {
	err := a.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router)
	logger.Fatal(err.Error())
}

func Start(a Application) {
	logger.Info("Starting banking-resource-api service")
	// Define routes and get a Router
	router := a.SetupRouter()

	host := ""
	port := ""
	if os.Getenv("RUN_AS_LAMBDA") != "true" {
		host = utils.CheckMandatoryEnvVar("API_HOST")
		port = utils.CheckMandatoryEnvVar("API_PORT")
	}
	a.ListenAndServeRoutes(router, host, port)
}

func NewDefaultApplication(
	listenAndServe types.HttpListenAndServe,
	customerHandler CustomerHandler,
) Application {
	return DefaultApplication{
		ListenAndServe:  listenAndServe,
		CustomerHandler: customerHandler,
	}
}
