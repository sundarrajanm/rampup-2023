package main

import (
	"banking-resource-api/app"
	"banking-resource-api/domain"
	"banking-resource-api/service"
	"net/http"
	"os"

	"github.com/akrylysov/algnhsa"
	"github.com/jmoiron/sqlx"
)

func StartLambda(addr string, handler http.Handler) error {
	algnhsa.ListenAndServe(handler, nil)
	return nil
}

func main() {
	ServerStart := http.ListenAndServe

	if os.Getenv("RUN_AS_LAMBDA") == "true" {
		ServerStart = StartLambda
	}

	// Hexagonal Architecture Dependency Injection In Action
	customerRepo := domain.NewCustomerRepoMySql(sqlx.Open)
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := app.NewCustomerHandler(customerService)

	// Start the application
	app.Start(
		app.NewDefaultApplication(
			ServerStart,
			customerHandler,
		),
	)
}
