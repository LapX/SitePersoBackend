package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
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
	router.HandleFunc("/auth/google/login", oauthGoogleLogin)
	router.HandleFunc("/auth/google/callback", oauthGoogleCallback)
	router.Use(mux.CORSMethodMiddleware(router))
	initDatabaseConnection()
	log.Println("[INFO] Api listening on port " + port)
	log.Fatal(http.ListenAndServe(port, handlers.CORS(originsOk)(router)))

}

func initDatabaseConnection() {
	dataSourceName := getDataSourceName()
	database, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := database.Query("select * from oauthgoogle")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(rows)
	}
}

func getDataSourceName() string {
	dataSourceName := os.Getenv("DATABASE_URL")
	if dataSourceName == "" {
		dataSourceName = "postgres://zdlzizzwvxwvco:433ef3b77effdf353aac89a130baed1735e34cb326dbbaf72730f533998ce7b4@ec2-174-129-238-192.compute-1.amazonaws.com:5432/d1srf5edqjvuqv"
	}
	return dataSourceName
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

func oauthGoogleLogin(response http.ResponseWriter, request *http.Request) {

}
func oauthGoogleCallback(response http.ResponseWriter, request *http.Request) {

}
