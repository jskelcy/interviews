package main

import (
	"fmt"
)

func main() {

	tests := [][]int{
		{1, 4, 6, 8, 3, 6, 9, 3, 7, 3, 2, 76, 8},
	}
	for _, test := range tests {
		fmt.Printf("%v", dutchFlag(test, 4))
	}
}

func dutchFlag(input []int, pivot int) []int {
	less, equal, larger := 0, 0, len(input)-1
	pValue := input[pivot]
	for equal < larger {
		if input[equal] < pValue {
			input[less], input[equal] = input[equal], input[less]
			less++
			equal++
			continue
		}
		if input[equal] == pValue {
			equal++
			continue
		}
		input[equal], input[larger] = input[larger], input[equal]
		larger--
	}

	return input
}
