package main

import (
	"banking-resource-api/app"
	"banking-resource-api/domain"
	"banking-resource-api/service"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func main() {
	// Hexagonal Architecture Dependency Injection
	customerRepo := domain.NewCustomerRepoMySql()
	customerService := service.NewCustomerService(customerRepo)
	customerHandler := app.NewCustomerHandler(customerService)

	// 3rd Party Dependency Injection
	listenAndServe := http.ListenAndServe
	openSql := sqlx.Open

	// Start the application
	app.Start(
		app.NewDefaultApplication(
			listenAndServe,
			openSql,
			customerHandler,
		),
	)
}
