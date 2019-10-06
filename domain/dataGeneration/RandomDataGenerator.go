package dataGeneration

import "math/rand"

type EarningsGraphArray struct {
	EarningsGraphData []EarningsGraphData
}

type EarningsGraphData struct {
	ID     int
	Tuples []Tuple
}

type Tuple struct {
	Quarter  int
	Earnings int
}

func GenerateRandomDataList(nbrElements int) EarningsGraphArray {
	var generatedData []EarningsGraphData
	for i := 0; i < nbrElements; i++ {
		const baseAmount int = 10000
		var generatedTuples []Tuple
		for j := 1; j < 5; j++ {
			generatedTuples = append(generatedTuples, Tuple{j, rand.Intn(15000) + baseAmount})
		}
		generatedData = append(generatedData, EarningsGraphData{i, generatedTuples})
	}
	return EarningsGraphArray{EarningsGraphData: generatedData}
}
