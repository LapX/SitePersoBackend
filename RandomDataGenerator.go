package main

import "math/rand"

type Data struct {
	ID     int
	Tuples []Tuple
}

type Tuple struct {
	Quarter  int
	Earnings int
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
