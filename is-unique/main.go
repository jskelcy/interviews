package main

import (
	"fmt"
	"sort"
)

func main() {
	list := []int{1, 542, 76, 34, 7, 432, 47, 53, 7}
	fmt.Printf("%t\n", isUnique(list))
}

func isUnique(list []int) bool {
	// binary sort O(NlogN)
	sort.Ints(list)
	for index, value := range list {
		// check right
		if index+1 < len(list) {
			if list[index+1] == value {
				return true
			}
		}
	}

	return false
}
