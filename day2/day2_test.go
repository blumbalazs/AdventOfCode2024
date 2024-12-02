package main

import "testing"

var example = [][]int{
	{7, 6, 4, 2, 1},
	{1, 2, 7, 8, 9},
	{9, 7, 6, 2, 1},
	{1, 3, 2, 4, 5},
	{8, 6, 4, 4, 1},
	{1, 3, 6, 7, 9},
}

func TestChallange1Example(t *testing.T) {
	result := CalculateNumberOfSafeReports(example)

	if result != 2 {
		t.Fail()
	}
}
func TestChallange2Example(t *testing.T) {
	result := CalculateNumberOfSafeReportsWithToleration(example)
	if result != 4 {
		t.Fail()
	}
}
