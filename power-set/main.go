package main

import (
	"fmt"
	"math"
)

func main() {
	testData := []int{1, 2, 3}
	for _, sub := range powerSet(testData) {
		fmt.Printf("%v\n", sub)
	}
}

func powerSet(list []int) [][]int {
	output := make([][]int, int(math.Pow(2, float64(len(list)))))
	for j := 1; j < len(output); j++ {
		uj := uint64(j)
		fmt.Printf("%b\n", uj)
		var counter int
		for uj != 0 {
			if (uj & 1) == 1 {
				output[j] = append(output[j], list[counter])
			}
			uj >>= 1
			counter++
		}
	}

	return output
}
