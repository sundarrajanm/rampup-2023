package app

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Customer struct {
	Name    string `json:"name"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

func greet(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Hello World!")
}

func getAllCustomers(rw http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{"Shankar", "Chennai", "560048"},
		{"Sundar", "Bangalore", "560048"},
		{"Shankya", "Trivandrum", "560048"},
	}
	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(customers)
}

func getCustomer(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Fprint(rw, vars["cust_id"])
}
