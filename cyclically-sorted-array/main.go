package main

func main() {
	input := []int{378, 478, 550, 631, 103, 203, 220, 234, 279, 368}

	println(findStart(input))
}

func findStart(input []int) int {
	left, right := 0, len(input)-1
	for left < right {
		mid := (left + right) / 2
		if input[mid] > input[right] {
			left = mid + 1
			continue
		}
		right = mid
	}

	return left
}
