package adt

type Stack struct {
	size   int
	values []string
}

func NewStack() *Stack {
	return &Stack{0, make([]string, 5)}
}

func (s *Stack) Empty() bool {
	return s.size == 0
}

func (s *Stack) Add(value string) {
	s.values[s.size] = value
	s.size++
}

func (s *Stack) Size() int {
	return s.size
}

func (s *Stack) Pop() string {
	s.size--
	return s.values[s.size]
}
