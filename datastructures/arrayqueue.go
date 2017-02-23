package datastructures

import "errors"

type queue struct {
	data []int
	size int
	front int
}

func NewQueue(defaultCapacity int) *queue {
	return &queue{
		data: make([]int, defaultCapacity),
	}
}

func (q *queue) Length() int {
	return q.size
}

func (q *queue) IsEmpty() bool {
	return q.size == 0
}

func (q *queue) Front() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is Empty")
	}
	return q.data[q.front], nil
}

func (q *queue) Enqueue(e int) {
	if q.size == len(q.data) {
		q.resize(q.size * 2)
	}
	next := (q.front + q.size) % len(q.data)
	q.data[next] = e
	q.size += 1
}

func (q *queue) resize(capacity int) {
	old := q.data
	q.data = make([]int, capacity)
	walk := q.front
	for k := 0; k < q.size; k++ {
		q.data[k] = old[walk]
		walk = (1 + walk) % len(old)
	}
	q.front = 0
}

func (q *queue) Dequeue() (int, error) {
	if q.IsEmpty() {
		return 0, errors.New("Queue is Empty")
	}
	result := q.data[q.front]
	q.data[q.front] = 0
	q.front = (1 + q.front) % len(q.data)
	q.size -= 1
	return result, nil
}
