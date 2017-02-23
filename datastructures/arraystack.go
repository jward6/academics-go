package datastructures

import "errors"

type stack struct {
	data  []int
}

func NewStack() *stack {
	return &stack{
		data: make([]int, 0),
	}
}

func (s *stack) Length() int {
	return len(s.data)
}

func (s *stack) IsEmpty() bool {
	return s.Length() == 0
}

func (s *stack) Top() (int, error) {
	if s.IsEmpty(){
		return 0, errors.New("Stack is empty")
	}
	return s.data[s.Length() - 1], nil
}

func (s *stack) Push(e int){
	s.data = append(s.data, e)
}

func (s *stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, errors.New("Stack is empty")
	}

	result := s.data[s.Length() - 1]
	s.data = s.data[:s.Length()-1]
	return result, nil
}

