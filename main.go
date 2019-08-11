package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
)

type Data struct {
	ID     int
	Tuples []Tuple
}

type Tuple struct {
	Quarter  int
	Earnings int
}

func main() {
	const port = ":8080"
	router := mux.NewRouter()
	originsOk := handlers.AllowedOrigins([]string{"http://localhost:3000"})
	router.HandleFunc("/", getRoot).Methods("GET")
	router.HandleFunc("/data", getData).Methods("GET")
	router.Use(mux.CORSMethodMiddleware(router))
	fmt.Println("Api listening on port " + port)
	http.ListenAndServe(port, handlers.CORS(originsOk)(router))
}

func getRoot(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode("The api works")
}

func getData(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(generateRandomDataList(rand.Intn(6) + 1))
}

func generateRandomDataList(nbrElements int) []Data {
	var generatedData []Data
	for i := 0; i < nbrElements; i++ {
		const baseAmount int = 10000
		var generatedTuples []Tuple
		for j := 1; j < 5; j++ {
			generatedTuples = append(generatedTuples, Tuple{j, rand.Intn(15000) + baseAmount})
		}
		generatedData = append(generatedData, Data{i, generatedTuples})
	}
	return generatedData
}
