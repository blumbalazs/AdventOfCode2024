package main

import "math"

func CalculateNumberOfSafeReports(reportsOfLevels [][]int) int {

	numberOfSafeReports := 0
	for _, levels := range reportsOfLevels {

		growing := levels[0] < levels[1]
		stillGood := true
		for i := 1; i < len(levels) && stillGood; i++ {
			isDirectionGood := (levels[i] > levels[i-1]) == growing
			delta := int(math.Abs(float64(levels[i] - levels[i-1])))
			if !(isDirectionGood && delta >= 1 && delta <= 3) {
				stillGood = false
			}
		}
		if stillGood {
			numberOfSafeReports++
		}
	}

	return numberOfSafeReports
}
