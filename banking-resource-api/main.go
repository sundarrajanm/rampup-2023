package main

import (
	"banking-resource-api/app"
	"net/http"

	"github.com/jmoiron/sqlx"
)

func main() {
	app.Start(
		app.NewDefaultApplication(
			http.ListenAndServe,
			sqlx.Open,
		),
	)
}
