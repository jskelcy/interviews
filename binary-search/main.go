package main

import (
	"fmt"
	"sort"
)

func main() {
	list := []int{11, 26, 13, 5, 87, 22, 49, 76, 8, 34, 3, 45, 79, 93}
	sort.Ints(list)

	fmt.Printf("%v\n", list)
	println(index(list, 49))
}
func index(list []int, target int) int {
	if len(list) == 1 {
		if target == list[0] {
			return 0
		}
		return -1
	}

	middle := len(list) / 2
	if target == list[middle] {
		return middle
	}
	if target < list[middle] {
		return index(list[:middle], target)
	}

	targetIndex := index(list[middle:], target)
	if targetIndex != -1 {
		return middle + targetIndex
	}

	return -1
}
