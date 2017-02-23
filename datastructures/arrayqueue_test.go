package datastructures

import "testing"

func TestEnqueue(t *testing.T) {
	q := NewQueue(10)
	q.Enqueue(10)
	q.Enqueue(15)
	result,err := q.Front()

	if err != nil {
		t.Error(err)
	}

	if result != 10 {
		t.Errorf("Front() == %q, want %q", result, 10)
	}
}

func TestDequeue(t *testing.T) {
	q := NewQueue(10)
	q.Enqueue(10)
	q.Enqueue(15)
	q.Enqueue(20)
	result,err := q.Dequeue()

	if err != nil {
		t.Error(err)
	}

	if result != 10 {
		t.Errorf("Dequeue() == %q, want %q", result, 10)
	}
}

func TestResize(t *testing.T) {
	q := NewQueue(3)
	q.Enqueue(10)
	q.Enqueue(15)
	q.Enqueue(20)
	q.Enqueue(25)

	front, err := q.Front()

	if err != nil {
		t.Error(err)
	}

	if q.Length() != 4 {
		t.Errorf("Length() == %q, want %q", q.Length(), 4)
	}

	if front != 10 {
		t.Errorf("Front() == %q, want %q", front, 10)
	}
}
