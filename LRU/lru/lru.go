package lru

import (
	"fmt"
)

type node struct {
	id   string
	data interface{}
	prev *node
	next *node
}

type LRU struct {
	lookup   map[string]*node
	listHead *node
	listTail *node
	maxSize  int
	currSize int
}

func NewLRU(maxSize int) (LRU, error) {
	if maxSize == 1 {
		return LRU{}, fmt.Errorf("maxSize must be larger than 1")
	}
	return LRU{
		lookup:   make(map[string]*node, maxSize),
		maxSize:  maxSize,
		currSize: 0,
	}, nil
}

func (l *LRU) Insert(id string, data interface{}) {
	if l.maxSize == l.currSize {
		// remove last node if the list is full
		tail := l.listTail
		l.listTail = l.listTail.prev
		l.listTail.next = nil

		delete(l.lookup, tail.id)
		l.currSize--
	}

	// Insert new node at front of list
	newNode := &node{
		id:   id,
		data: data,
	}
	if l.listHead == nil {
		l.listHead = newNode
		l.listTail = newNode
	} else {
		l.listHead.prev = newNode
		l.listHead = l.listHead.prev
	}
	l.lookup[id] = newNode
	l.currSize++
}

func (l *LRU) Read(id string) (interface{}, error) {
	readNode, exists := l.lookup[id]
	if !exists {
		return nil, fmt.Errorf("id: %s, does not exist", id)
	}

	// Just return if already at the head
	if readNode == l.listHead {
		return readNode.data, nil
	}
	// If node is at the tail, set the tail back one
	if readNode == l.listTail {
		l.listTail = l.listTail.prev
	}

	// remove node from list and add to front
	readNode.prev.next = readNode.next
	// If the node is the tail do not set next
	if readNode.next != nil {
		readNode.next.prev = readNode.prev
	}
	// set next to the current head and prev to nil
	readNode.next = l.listHead
	readNode.prev = nil
	// set current head prev to new head
	l.listHead.prev = readNode
	// reset head
	l.listHead = l.listHead.prev

	return readNode.data, nil
}
