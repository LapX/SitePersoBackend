package resource

import (
	"encoding/json"
	"github.com/LapX/SitePersoBackend/dataGeneration"
	"log"
	"math/rand"
	"net/http"
)

func GetData(response http.ResponseWriter, request *http.Request) {
	log.Println("[INFO] /data got called")
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(dataGeneration.GenerateRandomDataList(rand.Intn(6) + 1))
}
