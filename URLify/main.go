package main

import (
	"fmt"
)

func main() {
	input := make([]rune, 0, 13)
	for _, letter := range "Mr John Smith    " {
		input = append(input, letter)
	}

	fmtInput := urlify(input)

	fmt.Println(string(fmtInput))
}

func urlify(input []rune) []rune {
	for i, char := range input {
		if char == ' ' {
			copy(input[i+3:], input[i+1:])
			input[i] = '%'
			input[i+1] = '2'
			input[i+2] = '0'
		}
	}

	return input
}
