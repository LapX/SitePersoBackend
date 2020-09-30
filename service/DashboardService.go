package service

import (
	"github.com/LapX/SitePersoBackend/dependencies/database"
	"github.com/LapX/SitePersoBackend/domain/dataGeneration"
	"math/rand"
)

type Graphs struct {
	EarningsGraphArray []dataGeneration.EarningsGraphArray
}

func GetGraphs(token string) Graphs {
	var graphs Graphs
	var numberOfEarningsGraphs int
	if token != "undefined" {
		numberOfEarningsGraphs = database.FetchNumberOfEarningsGraphs(token)
	} else {
		numberOfEarningsGraphs = database.FetchNumberOfEarningsGraph()
	}
	for i := 0; i < numberOfEarningsGraphs; i++ {
		graphs.EarningsGraphArray = append(graphs.EarningsGraphArray, dataGeneration.GenerateRandomDataList(rand.Intn(6)+1))
	}

	return graphs
}

func AddGraphs(token string, amount int) {
	if amount > 20 {
		database.ModifyNumberOfGraphs(token, 20)
	} else {
		database.ModifyNumberOfGraphs(token, amount)
	}

}
