package stack

type node struct {
	value interface{}
	prev  *node
}

type Stack struct {
	top    *node
	length int
}

func New() *Stack {
	return &Stack{top: nil, length: 0}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Peek() interface{} {
	if s.Len() == 0 {
		return nil
	}
	return s.top.value
}

func (s *Stack) Pop() interface{} {
	if s.Len() == 0 {
		return nil
	}
	popped := s.top.value
	s.top = s.top.prev
	s.length--
	return popped
}

func (s *Stack) Push(value interface{}) {
	n := &node{value: value, prev: s.top}
	s.top = n
	s.length++
}
