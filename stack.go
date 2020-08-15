package stopka

type Stack struct {
	stack []Value
}

func (s *Stack) Push(element Value) {
	s.stack = append(s.stack, element)
}

func (s *Stack) Pop() (Value, bool) {
	if s.IsEmpty() {
		return nil, false
	}
	top := s.stack[len(s.stack)-1]
	s.stack = s.stack[0 : len(s.stack)-1]
	return top, true
}

func (s *Stack) IsEmpty() bool {
	return len(s.stack) == 0
}

func (s *Stack) Values() []Value {
	return s.stack
}
