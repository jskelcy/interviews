package intlist

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	head *Node
	tail *Node
	len  int
}

func MakeList(list []int) *Node {
	head := &Node{}
	currNode := head
	for i, val := range list {
		currNode.Value = val
		if i < len(list)-1 {
			currNode.Next = &Node{}
		}
		currNode = currNode.Next
	}

	return head
}

func NewList() *LinkedList {
	doughnut := &Node{}

	return &LinkedList{
		head: doughnut,
		tail: doughnut.Next,
		len:  1,
	}
}

func (l *LinkedList) Head() *Node {
	return l.head.Next
}

func (l *LinkedList) Tail() *Node {
	return l.tail
}

func (l *LinkedList) Len() int {
	return l.len
}

func (l *LinkedList) Insert(value int) {
	defer func() { l.len++ }()
	n := &Node{
		Value: value,
	}
	tail := l.Tail()
	if tail == nil {
		l.head.Next = n
		l.tail = n
		return
	}

	l.Tail().Next = n
	l.tail = l.tail.Next
}

func (l *LinkedList) InsertOrdered(value int) {
	defer func() { l.len++ }()
	n := &Node{
		Value: value,
	}

	p := l.Head()
	if p == nil {
		l.head.Next = n
		return
	}

	if value < p.Value {
		n.Next = p
		l.head.Next = n
		return
	}

	for p.Next != nil {
		if p.Value == value || (p.Value < value && value < p.Next.Value) {
			n.Next = p.Next
			p.Next = n
			return
		}
		p = p.Next
	}
	p.Next = n
}

func (l *LinkedList) ToSlice() []int {
	slice := make([]int, 0, l.Len())
	p := l.Head()

	for p != nil {
		slice = append(slice, p.Value)
		p = p.Next
	}
	return slice
}
