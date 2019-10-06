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
	numberOfEarningsGraphs := 6
	if token != "undefined" {
		numberOfEarningsGraphs = database.FetchNumberOfEarningsGraphs(token)
	}
	for i := 0; i < numberOfEarningsGraphs; i++ {
		graphs.EarningsGraphArray = append(graphs.EarningsGraphArray, dataGeneration.GenerateRandomDataList(rand.Intn(6)+1))
	}

	return graphs
}
