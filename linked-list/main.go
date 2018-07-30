package main

import (
	"fmt"
	intlist "interview/linked-list/int-list"
)

func main() {
	nums := []int{1, 3, 4, 45, 73, 54, 23, 134, 7, 365, 4, 2534, 13, 54, 76, 5, 413, 3, 45, 6, 35, 46, 7}

	list := intlist.NewList()
	for _, num := range nums {
		list.InsertOrdered(num)
	}

	slice := list.ToSlice()
	fmt.Printf("%v\n", slice)
}
