package main

import (
	"fmt"
	"math"
)

func main() {
	data := []int{1, 3, 2, 4, 4}
	fmt.Printf("%v\n", findPairs(data, 5))
	fmt.Printf("%v\n", findPairsFaster(data, 5))
}

func findPairs(n []int, target int) [][]int {
	indexMap := map[int][]int{}
	for index, value := range n {
		indexMap[value] = append(indexMap[value], index)
	}

	result := [][]int{}
	for key, value := range indexMap {
		difference := target - key
		diffIndex, ok := indexMap[difference]
		if !ok {
			continue
		}
		for _, currIndex := range value {
			for _, currDiffIndex := range diffIndex {
				result = append(result, []int{currIndex, currDiffIndex})
			}
		}
	}

	return result
}

// Only move through half of the array
func findPairsFaster(n []int, target int) [][]int {
	indexMap := make([][]int, target+1)
	for index, value := range n {
		indexMap[value] = append(indexMap[value], index)
	}

	result := [][]int{}
	half := int(math.Ceil(float64(len(indexMap)) / float64(2)))
	for i := 0; i < half; i++ {
		difference := target - i
		if len(indexMap[difference]) == 0 {
			continue
		}
		for _, currIndex := range indexMap[i] {
			for _, currDiffIndex := range indexMap[difference] {
				result = append(result, []int{currIndex, currDiffIndex})
			}
		}
	}

	return result
}
