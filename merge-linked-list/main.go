package main

import (
	"fmt"
	"interview/linked-list/int-list"
)

func main() {
	slices := [][]int{
		[]int{2, 5, 7},
		[]int{3, 11},
	}

	lists := make([]*intlist.LinkedList, len(slices))
	for i, slice := range slices {
		list := intlist.NewList()
		for _, value := range slice {
			list.InsertOrdered(value)
		}
		lists[i] = list
	}

	fmt.Println(merge(lists[0], lists[1]).ToSlice())
}

func merge(l1, l2 *intlist.LinkedList) *intlist.LinkedList {
	p1 := l1.Head()
	p2 := l2.Head()
	newList := intlist.NewList()

	for p1 != nil && p2 != nil {
		if p1.Value == p2.Value {
			newList.Insert(p1.Value)
			newList.Insert(p2.Value)
			p1 = p1.Next
			p2 = p2.Next
			continue
		}
		if p1.Value < p2.Value {
			newList.Insert(p1.Value)
			p1 = p1.Next
			continue
		}
		newList.Insert(p2.Value)
		p2 = p2.Next
	}

	if p1 == nil {
		for p2 != nil {
			newList.Insert(p2.Value)
			p2 = p2.Next
		}
	}

	if p2 == nil {
		for p1 != nil {
			newList.Insert(p1.Value)
			p1 = p1.Next
		}
	}

	return newList
}
