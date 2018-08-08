package stack

import "errors"

type Stack struct {
	items []rune
}

func NewStack() *Stack {
	return &Stack{
		items: []rune{},
	}
}

func (s *Stack) Peek() rune {
	return s.items[len(s.items)-1]
}

func (s *Stack) Pop() (rune, error) {
	if len(s.items) == 0 {
		return ' ', errors.New("can not pop from empty Stack")
	}

	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item, nil
}

func (s *Stack) Push(item rune) {
	s.items = append(s.items, item)
}
