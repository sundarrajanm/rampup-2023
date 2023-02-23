package main

import (
	"banking-resource-api/app"
	"net/http"
)

func main() {
	app.Start(
		app.NewDefaultApplication(
			http.ListenAndServe,
		),
	)
}
