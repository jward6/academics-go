package datastructures

import "testing"

func TestLength(t *testing.T) {
	s := NewStack()
	s.Push(10)
	if s.Length() != 1 {
		t.Errorf("Length() == %q, want %q", s.Length(), 1)
	}
}

func TestTop(t *testing.T) {
	s := NewStack()
	s.Push(10)
	result, err := s.Top()
	if err != nil {
		t.Fail()
	}
	if result != 10 {
		t.Errorf("Top() == %q, want %q", result, 10)
	}
}

func TestPush(t *testing.T) {
	s := NewStack()
	s.Push(10)
	s.Push(15)
	result, err := s.Top()
	if err != nil {
		t.Error(err)
	}
	if result != 15 {
		t.Errorf("Top() == %q, want %q", result, 15)
	}
	length := s.Length()
	if length != 2 {
		t.Errorf("Length() == %q, want %q", length, 2)
	}
}

func TestPop(t *testing.T) {
	s := NewStack()
	s.Push(10)
	s.Push(15)
	result, err := s.Pop()
	if err != nil {
		t.Error(err)
	}
	if result != 15 {
		t.Errorf("Pop() == %q, want %q", result, 15)
	}
}
