package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	fmt.Println("Starting backend app")

	router := mux.NewRouter()

	// define routes
	router.HandleFunc("/greet", greet)
	router.HandleFunc("/customers", getAllCustomers)
	router.HandleFunc("/customers/{cust_id}", getCustomer)

	// Start the server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
