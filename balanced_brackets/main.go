package main

import (
	"bufio"
	"interview/stack"
	"log"
	"os"
)

func main() {
	fileLocal := "./input.txt"
	inputFile, err := os.Open(fileLocal)
	if err != nil {
		log.Fatalf("unable to read file %s", fileLocal)
	}

	sc := bufio.NewScanner(inputFile)
	for sc.Scan() {
		if isBalanced(sc.Text()) {
			log.Print("YES")
		} else {
			log.Print("NO")
		}
	}
	if err := sc.Err(); err != nil {
		log.Fatalf("error scanner %s", fileLocal)
	}
}

func isBalanced(tower string) bool {
	s := stack.NewStack()
	for _, brick := range tower {
		if isOpener(brick) {
			s.Push(brick)
		} else {
			nextOpenerBrick, err := s.Pop()
			if err != nil {
				return false
			}
			expectedOpener := opener(brick)
			if expectedOpener != nextOpenerBrick {
				return false
			}
		}
	}

	return true
}

func isOpener(brick rune) bool {
	switch brick {
	case '(':
		return true
	case '[':
		return true
	case '{':
		return true
	default:
		return false
	}
}

func opener(brick rune) rune {
	switch brick {
	case ')':
		return '('
	case ']':
		return '['
	case '}':
		return '{'
	default:
		return '('
	}
}
