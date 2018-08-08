package main

import (
	"fmt"
)

func main() {
	testCase := []int{-14, -10, 2, 108, 108, 243, 285, 285, 285, 285, 285, 413}

	firstOccurence := findFirst(testCase, 285)
	fmt.Println(firstOccurence)
}

func findFirst(input []int, value int) int {
	left, right := 0, len(input)-1
	smallestFound := -1
	for left <= right {
		pivot := (right + left) / 2
		if input[pivot] < value {
			left = pivot + 1
			continue
		}
		if input[pivot] == value {
			smallestFound = pivot
		}
		right = pivot - 1
	}

	return smallestFound
}
