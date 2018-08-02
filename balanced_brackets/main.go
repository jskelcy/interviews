package main

import (
	"bufio"
	"errors"
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
	s := newStack()
	for _, brick := range tower {
		if isOpener(brick) {
			s.push(brick)
		} else {
			nextOpenerBrick, err := s.pop()
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

type stack struct {
	items []rune
}

func newStack() *stack {
	return &stack{
		items: []rune{},
	}
}

func (s *stack) peek() rune {
	return s.items[len(s.items)-1]
}

func (s *stack) pop() (rune, error) {
	if len(s.items) == 0 {
		return ' ', errors.New("can not pop from empty stack")
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func (s *stack) push(item rune) {
	s.items = append(s.items, item)
}
