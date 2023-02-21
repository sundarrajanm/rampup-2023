package main

import (
	"banking-resource-api/app"
	"banking-resource-api/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting banking resource service")

	// Define Customers Handler
	router := mux.NewRouter()
	ch := app.CustomerHandler{Service: service.NewCustomerService()}
	router.HandleFunc("/customers", ch.GetAllCustomers)

	// Start Server
	host := "localhost"
	port := "8000"
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), router))
}
