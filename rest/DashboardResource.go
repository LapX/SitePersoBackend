package rest

import (
	"encoding/json"
	"github.com/LapX/SitePersoBackend/domain/dataGeneration"
	"github.com/LapX/SitePersoBackend/service"
	"math/rand"
	"net/http"
)

func GetData(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(dataGeneration.GenerateRandomDataList(rand.Intn(6) + 1))
}

func GetGraphs(response http.ResponseWriter, request *http.Request) {
	token, err := request.URL.Query()["token"]
	response.Header().Set("Content-Type", "application/json")
	if err {
		graphs := service.GetGraphs(token[0])
		json.NewEncoder(response).Encode(graphs)
	} else {
		json.NewEncoder(response).Encode("token query param missing")
	}
}
