package main

import (
	"fmt"
	intlist "interview/linked-list/int-list"
)

func main() {
	slices := [][]int{
		{1, 4, 5}, {1, 3, 4}, {2, 6},
	}
	lists := make([]*intlist.Node, len(slices))
	for i, slice := range slices {
		list := intlist.NewList()
		for _, value := range slice {
			list.InsertOrdered(value)
		}
		lists[i] = list.Head()
	}

	merged := mergeKLists(lists)
	final := []int{}
	for merged != nil {
		final = append(final, merged.Value)
		merged = merged.Next
	}

	fmt.Printf("%v\n", final)
}

func mergeKLists(lists []*intlist.Node) *intlist.Node {
	result := &intlist.Node{}
	resultRunner := result
	for len(lists) != 0 {
		var smallestSeen int
		for i, list := range lists {
			if list.Value < lists[smallestSeen].Value {
				smallestSeen = i
			}
		}
		// add node to end of result
		resultRunner.Next = lists[smallestSeen]
		resultRunner = resultRunner.Next

		lists[smallestSeen] = lists[smallestSeen].Next
		// Now that resultsRunner and lists[smallestSeen] have been decoupled you can
		// remove change the next to nil.
		resultRunner.Next = nil

		//remove empty list
		if lists[smallestSeen] == nil {
			copy(lists[smallestSeen:], lists[smallestSeen+1:])
			lists = lists[:len(lists)-1]
		}

	}
	result = result.Next
	return result
}
