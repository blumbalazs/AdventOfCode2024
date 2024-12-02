package main

import (
	"testing"
)

var exampleLeftSide = []int{3, 4, 2, 1, 3, 3}
var exampleRightSide = []int{4, 3, 5, 3, 9, 3}

func TestChallange1Exmaple(t *testing.T) {
	result := GetSumOfParedListDelta(exampleLeftSide, exampleRightSide)

	if result != 11 {
		t.Fail()
	}
}

func TestChallange2Example(t *testing.T) {
	result := CalculateSimilarityScore(exampleLeftSide, exampleRightSide)

	if result != 31 {
		t.Fail()
	}
}
