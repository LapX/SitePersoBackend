package rest

import (
	"encoding/json"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

func InitServer() {
	port := getPort()
	router := mux.NewRouter()
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000", "https://lapx.github.io"})
	router.HandleFunc("/", getRoot).Methods("GET")
	router.HandleFunc("/data", getData).Methods("GET")
	router.HandleFunc("/dataGraphs", getGraphs).Methods("GET")
	router.HandleFunc("/auth/google/login", oauthGoogleLogin).Methods("GET")
	router.HandleFunc("/auth/google/callback", oauthGoogleCallback).Methods("GET")
	router.HandleFunc("/auth", getUserInfo).Methods("GET")
	router.HandleFunc("/addGraphs", modifyAmountOfGraphs).Methods("POST")
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
