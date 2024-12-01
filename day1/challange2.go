package main

func CalculateSimilarityScore(leftSideList []int, rightSideList []int) int {

	leftSideOccurance := countValueOccurance(leftSideList)

	similarityScore := 0
	for _, value := range rightSideList {
		similarityScore += value * leftSideOccurance[value]
	}

	return similarityScore
}

func countValueOccurance(list []int) map[int]int {
	occurance := map[int]int{}

	for _, value := range list {
		occurance[value] = occurance[value] + 1
	}

	return occurance
}
