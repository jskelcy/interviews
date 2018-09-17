package main

import intlist "interview/linked-list/int-list"

func main() {
	head := intlist.MakeList([]int{3, 5, 8, 5, 10, 2, 1})
	partitionList(head, 8)

	printRunner := head
	for printRunner != nil {
		println(printRunner.Value)
		printRunner = printRunner.Next
	}
}

func partitionList(head *intlist.Node, pivot int) {
	gtEtStart := head
	runner := head
	var prev *intlist.Node

	for gtEtStart.Next != nil {
		gtEtStart = gtEtStart.Next
	}
	tail := gtEtStart

	for runner != gtEtStart.Next {
		if runner.Value < pivot {
			prev = runner
			runner = runner.Next
		}

		if prev != nil {
			// remove node
			prev.Next = runner.Next
			swapRunner := runner
			runner = swapRunner.Next
			swapRunner.Next = nil
			// add to end
			tail.Next = swapRunner
			tail = tail.Next
		} else {
			// advance head
			head = head.Next
			// Remove runner
			runner = runner.Next
			swapRunner := runner
			swapRunner.Next = nil
			// add runner to tail
			tail.Next = swapRunner
			tail = tail.Next
		}
	}
}
