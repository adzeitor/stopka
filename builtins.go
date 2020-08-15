package stopka

// swap ( a b -- b a )
func (m *Machine) Swap() {
	first := m.Pop()
	second := m.Pop()
	m.Push(first)
	m.Push(second)
}

// dup ( a -- a a )
func (m *Machine) Dup() {
	top := m.Pop()
	m.Push(top)
	m.Push(top)
}

// drop ( a -- )
func (m *Machine) Drop() {
	_ = m.Pop()
}

func (m *Machine) Plus() {
	first := m.Pop()
	second := m.Pop()
	if m.IsHalted() {
		return
	}
	result := first.Plus(second)
	m.Push(result)
}

func (m *Machine) Minus() {
	first := m.Pop()
	second := m.Pop()
	if m.IsHalted() {
		return
	}

	result := first.Minus(second)
	m.Push(result)
}

func (m *Machine) ToString() {
	top := m.Pop()
	if m.IsHalted() {
		return
	}

	result := top.ToString()
	m.Push(result)
}

// map ( func list -- list' )
func (m *Machine) Map() {
	// FIXME: panic
	fn := m.Pop().(List)
	list := m.Pop().(List)
	if m.IsHalted() {
		return
	}

	result := make(List, 0, len(list))
	// save stack
	stack := m.stack

	m.stack = &Stack{}
	for _, x := range list {
		m.Push(x)
		for _, operator := range fn {
			m.evalValue(operator)
		}
		result = append(result, m.Pop())
	}
	// restore stack
	m.stack = stack
	if m.IsHalted() {
		return
	}
	m.stack.Push(result)
}

// apply ( (func|list) -- result )
func (m *Machine) Apply() {
	fn := m.Pop()
	if m.IsHalted() {
		return
	}

	m.evalValue(Identifier(fn.(Symbol)))
}
