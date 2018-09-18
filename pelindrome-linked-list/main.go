package main

import (
	"fmt"
	intlist "interview/linked-list/int-list"
)

func main() {
	listEven := intlist.MakeList([]int{1, 2, 3, 4, 4, 3, 2, 1})
	listOdd := intlist.MakeList([]int{1, 2, 3, 4, 5, 4, 3, 2, 1})
	listBad := intlist.MakeList([]int{1, 8, 3, 4, 5, 4, 3, 2, 1})
	fmt.Println(palindromeLinkedList(listEven))
	fmt.Println(palindromeLinkedList(listOdd))
	fmt.Println(palindromeLinkedList(listBad))
}

func palindromeLinkedList(head *intlist.Node) bool {
	fastRunner := head
	slowRunner := head
	stack := []int{}

	for fastRunner != nil && fastRunner.Next != nil {
		stack = append(stack, slowRunner.Value)
		slowRunner = slowRunner.Next
		fastRunner = fastRunner.Next.Next
	}

	if fastRunner != nil {
		slowRunner = slowRunner.Next
	}

	for i := len(stack) - 1; slowRunner != nil; i-- {
		if stack[i] != slowRunner.Value {
			return false
		}
		slowRunner = slowRunner.Next
	}

	return true
}
