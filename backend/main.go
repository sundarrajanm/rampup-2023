package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name string `json:"name"`
	City string `json:"city"`
	ZipCode string `json:"zip_code"`
}

func main() {

	// define routes
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", getAllCustomers)

	// Start the server
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func greet(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello World!")
}

func getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	customers := []Customer {
		{ "Hello", "Bangalore", "560048" },
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(customers)
}
