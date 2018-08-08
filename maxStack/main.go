package main

import "interview/maxStack/maxstack"

func main() {
	nums := []int{1, 5, 1323, 654, 34, 5, 86, 45, 7, 35, 54}

	maxStack := maxstack.NewMaxStack()

	for _, num := range nums {
		maxStack.Push(num)
	}
	println(maxStack.Max())
	println(maxStack.Pop())
	println(maxStack.Pop())
	println(maxStack.Pop())
	println(maxStack.Pop())
	println(maxStack.Pop())
}
