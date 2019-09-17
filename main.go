package main

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"os"
)

func main() {
	port := getPort()
	router := mux.NewRouter()
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000", "https://lapx.github.io"})
	router.HandleFunc("/", getRoot).Methods("GET")
	router.HandleFunc("/data", getData).Methods("GET")
	router.Use(mux.CORSMethodMiddleware(router))
	log.Println("[INFO] Api listening on port " + port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(originsOk)(router)))
}

func getPort() string {
	port := ":" + os.Getenv("PORT")
	if port == ":" {
		port = ":8080"
	}
	return port
}

func getRoot(response http.ResponseWriter, request *http.Request) {
	log.Println("[INFO] / got called")
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode("The api works")
}

func getData(response http.ResponseWriter, request *http.Request) {
	log.Println("[INFO] /data got called")
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(generateRandomDataList(rand.Intn(6) + 1))
}
