package main

import (
	"math"
	"sort"
)

func GetSumOfParedListDelta(leftSideList []int, rightSideList []int) int {

	sort.Ints(leftSideList)
	sort.Ints(rightSideList)

	deltaSum := 0
	for i, value := range leftSideList {
		deltaSum += int(math.Abs(float64(rightSideList[i]) - float64(value)))
	}
	return deltaSum
}
