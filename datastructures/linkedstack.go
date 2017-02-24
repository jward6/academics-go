package datastructures

import "errors"

var EmptyErr = errors.New("Stack is empty")

type node struct {
	element int
	next *node
}

type linkedStack struct {
	header *node
	size  int
}

func NewLinkedStack() *linkedStack {
	return &linkedStack {
		header: &node {},
		size: 0,
	}
}

func (l *linkedStack) Length() int {
	return l.size
}

func (l *linkedStack) IsEmpty() bool {
	return l.Length() == 0
}

func (l *linkedStack) Top() (int, error) {
	if l.IsEmpty() {
		return 0, EmptyErr
	}
	return l.header.next.element, nil
}

func (l *linkedStack) Pop() (int, error) {
	if l.IsEmpty() {
		return 0, EmptyErr
	}
	result := l.header.next
	l.header.next = result.next
	l.size -= 1
	return result.element, nil
}

func (l *linkedStack) Push(e int) {
	node := &node{
		element: e,
		next: l.header.next,
	}
	l.header.next = node
	l.size += 1
}

