package maxstack

import "container/heap"

type element struct {
	last  *element
	next  *element
	val   int
	index int // The index of the item in the heap.
}

type MaxStack struct {
	head    *element
	tail    *element
	maxHeap *priorityQueue
}

func NewMaxStack() MaxStack {
	maxHeap := &priorityQueue{}
	heap.Init(maxHeap)
	return MaxStack{
		maxHeap: maxHeap,
	}
}

func (m *MaxStack) Push(val int) {
	ele := &element{
		last: m.tail,
		val:  val,
	}
	if m.tail == nil {
		//empty
		m.head = ele
		m.tail = ele
	} else {
		m.tail.next = ele
		m.tail = ele
	}
	heap.Push(m.maxHeap, ele)
}

func (m *MaxStack) Pop() int {
	if len(*m.maxHeap) == 0 {
		return 0
	}
	ele := m.head
	m.head = ele.next
	if m.head == nil {
		m.tail = nil
	}
	max := m.maxHeap.peek()
	maxEle := max.(*element)
	if maxEle == ele {
		heap.Pop(m.maxHeap)
	}

	return ele.val
}

func (m *MaxStack) Max() int {
	if len(*m.maxHeap) != 0 {
		max := heap.Pop(m.maxHeap)
		maxEle := max.(*element)
		if maxEle.last != nil {
			maxEle.last.next = maxEle.next
		}
		if maxEle.next != nil {
			maxEle.next.last = maxEle.last
		}
		maxEle.next = nil
		maxEle.last = nil
		return maxEle.val
	}

	return 0
}

// A priorityQueue implements heap.Interface and holds items.
type priorityQueue []*element

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].val > pq[j].val
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *priorityQueue) Push(x interface{}) {
	n := len(*pq)
	ele := x.(*element)
	ele.index = n
	*pq = append(*pq, ele)
}

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	ele := old[n-1]
	ele.index = -1 // for safety
	*pq = old[0 : n-1]
	return ele
}

// update modifies the priority and value of an item in the queue.
func (pq *priorityQueue) update(ele *element, val int) {
	ele.val = val
	heap.Fix(pq, ele.index)
}

func (pq *priorityQueue) peek() interface{} {
	peekable := *pq
	return peekable[len(peekable)-1]
}
