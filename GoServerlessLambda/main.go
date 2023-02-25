package main

import (
	"encoding/json"
	"net/http"

	"github.com/akrylysov/algnhsa"
	"github.com/gorilla/mux"
)

//	func GenerateResponse(Body string, Code int) events.APIGatewayProxyResponse {
//		return events.APIGatewayProxyResponse{Body: Body, StatusCode: Code}
//	}
//
//	func HandleRequest(_ context.Context, request events.LambdaFunctionURLRequest) (events.APIGatewayProxyResponse, error) {
//		return GenerateResponse("Hello World", 200), nil
//	}

func writeResponse(w http.ResponseWriter, code int, data interface{}) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(data)
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		writeResponse(w, http.StatusOK, [1]string{"Reached Home, Congratulations"})
	})

	router.HandleFunc("/customers", func(w http.ResponseWriter, r *http.Request) {
		writeResponse(w, http.StatusOK, [1]string{"Customers Home, Congratulations"})
	})

	algnhsa.ListenAndServe(router, nil)
	// lambda.Start(HandleRequest)
}
