package main

import (
	"banking-resource-api/app"
	"banking-resource-api/domain"
	"banking-resource-api/service"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func main() {
	// Hexagonal Architecture Dependency Injection In Action
	customerRepo := domain.NewCustomerRepoMySql(sqlx.Open)
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := app.NewCustomerHandler(customerService)

	// Start the application
	app.Start(
		app.NewDefaultApplication(
			http.ListenAndServe, // 3rd party dependency injection
			customerHandler,
		),
	)
}
