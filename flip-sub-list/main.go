package main

import (
	"fmt"
	"interview/linked-list/int-list"
)

func main() {
	nums := []int{11, 3, 5, 7, 2}
	list := intlist.NewList()
	for _, num := range nums {
		list.Insert(num)
	}

	fmt.Println(list.ToSlice())
	flipSubList(list, 1, 4)
	fmt.Println(list.ToSlice())
}

func flipSubList(list *intlist.LinkedList, start, stop int) {
	var (
		pre  *intlist.Node
		post *intlist.Node
		head *intlist.Node
		tail *intlist.Node
	)
	p := list.Head()
	pIndex := 0

	for pIndex < stop-1 {
		if pIndex == start-1 {
			pre = p
		}
		if pIndex == start {
			head = p
		}
		p = p.Next
		pIndex++
	}
	tail = p
	post = p.Next
	startFlip := head

	flip(startFlip, (stop-start)-1)
	pre.Next = tail
	head.Next = post
}

func flip(start *intlist.Node, totalFlips int) {
	flips := 0
	f1 := start
	f2 := start.Next
	for flips < totalFlips {
		stone := f2.Next
		f2.Next = f1
		f1 = f2
		f2 = stone
		flips++
	}
}
