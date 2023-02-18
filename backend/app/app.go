package app

import (
	"backend/domain"
	"backend/logger"
	"backend/service"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

const (
	API_ADDR  string = "API_ADDR"
	API_PORT         = "API_PORT"
	DB_USER          = "DB_USER"
	DB_PASSWD        = "DB_PASSWD"
	DB_ADDR          = "DB_ADDR"
	DB_PORT          = "DB_PORT"
	DB_NAME          = "DB_NAME"
)

func checkMandatoryEnvVars(vars ...string) {
	for _, v := range vars {
		if os.Getenv(v) == "" {
			errMsg := "Env variable " + v + " not found"
			logger.Error(errMsg)
			panic(errMsg)
		}
	}
}

func Start() {
	logger.Info("Starting backend app")
	checkMandatoryEnvVars(
		API_ADDR, API_PORT, DB_USER, DB_PASSWD, DB_ADDR, DB_PORT, DB_NAME,
	)

	router := mux.NewRouter()

	// Wiring
	// ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryDB())}

	// define routes``
	router.HandleFunc("/customers", ch.getAllCustomers)
	router.HandleFunc("/customers/{id:[0-9]+}", ch.getCustomer)

	// Start the server
	host := os.Getenv(API_ADDR)
	port := os.Getenv(API_PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router))
}
