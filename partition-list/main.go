package main

import (
	"fmt"
	intlist "interview/linked-list/int-list"
)

func main() {
	head := intlist.MakeList([]int{3, 5, 8, 5, 10, 2, 1})
	head = partitionList(head, 5)

	print(head)
}

func print(head *intlist.Node) {
	printString := []int{}
	printRunner := head
	for printRunner != nil {
		printString = append(printString, printRunner.Value)
		printRunner = printRunner.Next
	}

	fmt.Printf("%v\n", printString)
}

// move through the list. When you find a value >= to the pivot move it to the end of the list.
func partitionList(head *intlist.Node, pivot int) *intlist.Node {
	var times int
	tail := head
	runner := head
	var prev *intlist.Node

	for tail.Next != nil {
		tail = tail.Next
		times++
	}
	// return for empty list
	if times == 0 {
		return head
	}

	for i := 1; i <= times+1; i++ {
		// if value is less then pivot just move runner along
		if runner.Value <= pivot {
			prev = runner
			runner = runner.Next
			continue
		}

		swapRunner := runner
		runner = swapRunner.Next
		swapRunner.Next = nil
		if prev == nil {
			// advance head. There is no pervious.
			head = head.Next
		} else {
			// prev now skips swapRunner.
			prev.Next = runner
		}
		// add runner to tail
		tail.Next = swapRunner
		tail = tail.Next
	}

	return head
}
