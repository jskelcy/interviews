package main

import (
	"fmt"
)

const (
	right = "right"
	down  = "down"
	left  = "left"
	up    = "up"
)

func main() {
	image := [][]int{
		[]int{1, 2, 3, 4},
		[]int{12, 13, 14, 5},
		[]int{11, 16, 15, 6},
		[]int{10, 9, 8, 7},
	}

	orderedImage := spiral(image)
	for _, i := range orderedImage {
		fmt.Printf("%d, ", i)
	}
}

func spiral(image [][]int) []int {
	result := make([]int, 0, 16)
	x := 0
	y := 0
	direction := right
	layer := 1

	for len(result) < 16 {
		for i := 1; i <= len(image)-layer; i++ {
			fmt.Printf("Y: %d, X: %d, layer: %d\n", y, x, layer)
			result = append(result, image[y][x])
			move(&x, &y, direction, i == len(image)-layer)
		}
		direction, layer = turn(direction, layer)
	}

	return result
}

func move(x, y *int, direction string, stepIn bool) {
	switch direction {
	case right:
		*x++
	case down:
		*y++
	case left:
		*x--
	case up:
		if stepIn {
			*x++
			break
		}
		*y--
	}
}

func turn(direction string, layer int) (string, int) {
	var newDir string
	switch direction {
	case right:
		newDir = down
	case down:
		newDir = left
	case left:
		newDir = up
	case up:
		newDir = right
		layer = layer + 2
	}

	return newDir, layer
}
