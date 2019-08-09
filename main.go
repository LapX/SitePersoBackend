package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	const port = ":8081"
	router := mux.NewRouter()

	router.HandleFunc("/", getRoot).Methods("GET")
	fmt.Println("Api listening on port " + port)
	http.ListenAndServe(port, router)
}

func getRoot(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode("The api works")
}