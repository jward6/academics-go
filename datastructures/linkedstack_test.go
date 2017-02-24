package datastructures

import "testing"

func TestLinkedStackPush(t *testing.T) {
	l := NewLinkedStack()
	l.Push(10)
	l.Push(15)
	top,err := l.Top()
	
	if err != nil {
		t.Error(err)
	}

	length := l.Length()
	if length != 2 {
		t.Errorf("Length() == %q, want %q", length, 2)
	}

	if top != 15 {
		t.Errorf("Top() == %q, want %q", top, 15)
	}
}

func TestLinkedStackPop(t *testing.T) {
	l := NewLinkedStack()
	l.Push(10)
	l.Push(15)
	result, err := l.Pop()

	if err != nil {
		t.Error(err)
	}

	if result != 15 {
		t.Errorf("Pop() == %q, want %q", result, 15)
	}

	length := l.Length()

	if length != 1 {
		t.Errorf("Length() == %q, want %q", length, 1)
	}
}
