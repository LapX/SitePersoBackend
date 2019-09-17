package main

import (
	"testing"
)

func Test_givenNumberOfElements_whenOneElementIsAsked_thenReturnsOneElement(t *testing.T) {
	nbrElements := 1
	data := generateRandomDataList(nbrElements)
	if len(data) != nbrElements {
		t.Errorf("Data array is incorrect, contains more than %d elements", nbrElements)
	}
}

func Test_givenNumberOfElements_whenOneElementIsAsked_thenEarningsAreWithinBounds(t *testing.T) {
	nbrElements := 1
	lowerBound := 10000
	upperBound := 25000
	data := generateRandomDataList(nbrElements)
	for i := 0; i < 4; i++ {
		earnings := data[0].Tuples[i].Earnings
		if earnings < lowerBound || earnings > upperBound {
			t.Errorf("Earnings value is out of bounds Expected : %d <= earnings <= %d, Actual : %d", lowerBound, upperBound, earnings)
		}

	}
}
