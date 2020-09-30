package rest

import (
	"encoding/json"
	"github.com/LapX/SitePersoBackend/domain/dataGeneration"
	"github.com/LapX/SitePersoBackend/service"
	"math/rand"
	"net/http"
)

type TokenAmount struct {
	Token  string
	Amount int
}

func getData(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(dataGeneration.GenerateRandomDataList(rand.Intn(6) + 1))
}

func getGraphs(response http.ResponseWriter, request *http.Request) {
	token, err := request.URL.Query()["token"]
	response.Header().Set("Content-Type", "application/json")
	if err {
		graphs := service.GetGraphs(token[0])
		json.NewEncoder(response).Encode(graphs)
	} else {
		graphs := service.GetGraphs("undefined")
		json.NewEncoder(response).Encode(graphs)
	}
}

func modifyAmountOfGraphs(response http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)
	var tokenAmount TokenAmount
	err := decoder.Decode(&tokenAmount)
	if err != nil {
		panic(err)
	}
	service.AddGraphs(tokenAmount.Token, tokenAmount.Amount)
}
