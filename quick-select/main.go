package main

import (
	"fmt"
)

func main() {
	list := []int{2, 52, 74, 13, 10, 6, 7, 22}
	pivotLocal := quickSelect(list)
	fmt.Printf("%v\n", list)
	fmt.Printf("%v\n", pivotLocal)
}

func quickSelect(input []int) int {
	pivotInex := len(input) / 2
	pivot := input[pivotInex]
	input = append(input[:pivotInex], input[pivotInex+1:]...)
	input = append(input, pivot)

	lptr := 0
	rptr := len(input) - 2
	fmt.Printf("pivot %v\n", pivot)
	for lptr <= rptr {
		for input[lptr] < pivot {
			lptr++
		}
		for input[rptr] > pivot {
			rptr--
		}

		if lptr <= rptr {
			input[lptr], input[rptr] = input[rptr], input[lptr]
			lptr++
			rptr--
		}
	}

	input[lptr], input[len(input)-1] = input[len(input)-1], input[lptr]
	return lptr
}
