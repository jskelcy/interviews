package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	seed := []int{1, 5, 3, 6, 7, 4, 2, 4, 56, 7}
	copySeed := make([]int, len(seed))
	copy(copySeed, seed)

	fmt.Printf("%v\n", niaveRandomSubset(seed, 4))
	fmt.Printf("%v\n", inPlaceRandomSubset(copySeed, 4))
	fmt.Printf("%v\n", seed)
}

func niaveRandomSubset(input []int, size int) []int {
	output := make([]int, 0, size)
	rand.Seed(time.Now().Unix())
	for len(output) != cap(output) {
		index := rand.Intn(len(input))
		output = append(output, index)
	}

	return output
}

func inPlaceRandomSubset(i []int, size int) []int {
	input := i
	rand.Seed(time.Now().Unix())
	for i := 0; i < size; i++ {
		index := rand.Int() % len(input)
		input[i], input[index] = input[index], input[i]
	}

	return input[:size]
}
