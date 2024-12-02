package main

import "math"

func CalculateNumberOfSafeReportsWithToleration(reportsOfLevels [][]int) int {
	numberOfSafeReports := 0
	for _, levels := range reportsOfLevels {

		if areLevelsGoodWithError(levels) {
			numberOfSafeReports++
		}
	}

	return numberOfSafeReports
}

func areLevelsGoodWithError(levels []int) bool {
	isGood, index := isLevelsGood(levels)
	if !isGood {
		result, _ := isLevelsGood(RemoveIndex(levels, index))
		if result {
			return true
		}
		result2, _ := isLevelsGood(RemoveIndex(levels, index+1))
		if result2 {
			return true
		}

		result3, _ := isLevelsGood(RemoveIndex(levels, 0))
		return result3
	}
	return true
}

func isLevelsGood(levels []int) (isGood bool, index int) {
	growing := levels[0] < levels[1]
	for i := 0; i < len(levels)-1; i++ {
		isDirectionGood := (levels[i+1] > levels[i]) == growing
		delta := int(math.Abs(float64(levels[i+1] - levels[i])))
		if !(isDirectionGood && delta >= 1 && delta <= 3) {
			return false, i
		}
	}
	return true, 0
}

func RemoveIndex(s []int, index int) []int {
	ret := make([]int, 0)
	ret = append(ret, s[:index]...)
	return append(ret, s[index+1:]...)
}
