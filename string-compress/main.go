package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	testData := "aabcccccaaa"
	testData2 := "abcwyyda"
	fmt.Println(compress(testData))
	fmt.Println(compress(testData2))
}

func compress(data string) string {
	currentChar := string(data[0])
	currentCharCount := 0
	compressed := []string{}

	for _, rChar := range data {
		char := string(rChar)
		if char != currentChar {
			compressed = append(compressed, currentChar, strconv.Itoa(currentCharCount))
			currentChar = char
			currentCharCount = 0
		}
		currentCharCount++
	}
	compressed = append(compressed, currentChar, strconv.Itoa(currentCharCount))

	compressedString := strings.Join(compressed, "")
	if len(compressedString) > len(data) {
		return data
	}
	return compressedString
}
