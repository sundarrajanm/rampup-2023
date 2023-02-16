package app

import (
	"backend/domain"
	"backend/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Start() {
	fmt.Println("Starting backend app")

	router := mux.NewRouter()

	// Wiring
	ch := CustomerHandlers{service.NewCustomerService(domain.NewCustomerRepositoryStub())}

	// define routes
	router.HandleFunc("/customers", ch.getAllCustomers)

	// Start the server
	log.Fatal(http.ListenAndServe("localhost:8000", router))
}
