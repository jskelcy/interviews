package main

import (
	"fmt"
)

func main() {
	matrix := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	rotate(matrix)

	for _, row := range matrix {
		fmt.Printf("%v\n", row)
	}
}

func rotate(input [][]int) {
	for layer := 0; layer < len(input)/2; layer++ {
		first := layer
		last := len(input) - 1 - layer
		for i := first; i < last; i++ {
			offset := i - first
			top := input[first][i]

			// left to top
			input[first][i] = input[last-offset][first]
			// bottom to left
			input[last-offset][first] = input[last][last-offset]
			// right to bottom
			input[last][last-offset] = input[i][last]
			// top to right
			input[i][last] = top
		}
	}
}
