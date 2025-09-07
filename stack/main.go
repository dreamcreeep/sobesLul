package main

// Стек с возможностью получения минимального элемента
type Node struct {
	Value int
	Next  *Node
}

type MinStack struct {
	main *Node // основной стек
	min  *Node // стек минимумов
}

func NewMinStack() *MinStack {
	return &MinStack{}
}

func (s *MinStack) Push(val int) {
	s.main = &Node{Value: val, Next: s.main}
	if s.min == nil || val <= s.min.Value {
		s.min = &Node{Value: val, Next: s.min}
	}
}

func (s *MinStack) Pop() (int, bool) {
	if s.main == nil {
		return 0, false
	}
	val := s.main.Value
	s.main = s.main.Next

	if val == s.min.Value {
		s.min = s.min.Next
	}

	return val, true
}

func (s *MinStack) Top() (int, bool) {
	if s.main == nil {
		return 0, false
	}
	return s.main.Value, true
}

func (s *MinStack) GetMin() (int, bool) {
	if s.min == nil {
		return 0, false
	}
	return s.min.Value, true
}
