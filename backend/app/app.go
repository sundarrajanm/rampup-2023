package app

import (
	"backend/domain"
	"backend/logger"
	"backend/service"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
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
	dbClient := getDBClient()
	customerRepositoryDB := domain.NewCustomerRepositoryDB(dbClient)
	accountRepositoryDB := domain.NewAccountRepositoryDB(dbClient)
	ch := CustomerHandlers{service.NewCustomerService(customerRepositoryDB)}
	ah := AccountHandlers{service.NewAccountService(accountRepositoryDB)}

	// define routes``
	router.HandleFunc("/customers", ch.getAllCustomers)
	router.HandleFunc("/customers/{customer_id:[0-9]+}", ch.getCustomer)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account", ah.NewAccount).Methods(http.MethodPost)
	router.HandleFunc("/customers/{customer_id:[0-9]+}/account/{account_id:[0-9]+}", ah.MakeTransaction).Methods(http.MethodPost)

	// Start the server
	host := os.Getenv(API_ADDR)
	port := os.Getenv(API_PORT)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router))
}

func getDBClient() *sqlx.DB {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWD")
	host := os.Getenv("DB_ADDR")
	port := os.Getenv("DB_PORT")
	db := os.Getenv("DB_NAME")
	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, db)
	client, err := sqlx.Open("mysql", connectionString)
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)
	return client
}
