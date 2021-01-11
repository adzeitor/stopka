package stopka

import "errors"

type Values []Value

type Machine struct {
	operators []Value
	stack     *Stack
	words     map[string]Value
	Err       error
}

func New() *Machine {
	m := &Machine{
		stack: &Stack{},
		words: make(map[string]Value),
	}
	m.AddBuiltin("dup", m.Dup)
	m.AddBuiltin("swap", m.Swap)
	m.AddBuiltin("drop", m.Drop)
	m.AddBuiltin("+", m.Plus)
	m.AddBuiltin("-", m.Minus)
	m.AddBuiltin("map", m.Map)
	m.AddBuiltin("apply", m.Apply)
	m.AddBuiltin("string", m.ToString)
	m.AddBuiltin("define", m.Define)
	return m
}

func (m *Machine) evalValue(value Value) {
	switch v := value.(type) {
	case Builtin:
		v.fn()
	case Identifier:
		word, found := m.words[string(v)]
		if !found {
			m.ThrowException("unknown identifier " + string(v))
		}
		m.evalValue(word)
	default:
		m.Push(value)
	}
}

func (m *Machine) Eval(line string) *Machine {
	m.Err = nil
	values, err := parse(line)
	if err != nil {
		m.ThrowException(err.Error())
		return m
	}
	for _, value := range values {
		m.evalValue(value)
	}
	return m
}

func (m *Machine) AddWord(name string, value Value) *Machine {
	m.words[name] = value
	return m
}

func (m *Machine) AddBuiltin(name string, fn func()) *Machine {
	return m.AddWord(name, Builtin{name: name, fn: fn})
}

func (m *Machine) LoadOperators(line string) *Machine {
	operators, err := parse(line)
	if err != nil {
		m.ThrowException(err.Error())
		return m
	}
	m.operators = append(m.operators, operators...)
	return m
}

func (m *Machine) ThrowException(err string) *Machine {
	m.Err = errors.New(err)
	return m
}

func (m *Machine) IsHalted() bool {
	return m.Err != nil
}

func (m *Machine) Push(element Value) {
	if e, isException := element.(Exception); isException {
		m.ThrowException(string(e))
	}
	if m.IsHalted() {
		return
	}
	m.stack.Push(element)
}

func (m *Machine) Stack() []Value {
	return m.stack.Values()
}

func (m *Machine) Operators() []Value {
	return m.operators
}

func (m *Machine) Pop() Value {
	if m.IsHalted() {
		return nil
	}

	value, ok := m.stack.Pop()
	if !ok {
		m.ThrowException("stack is empty")
		return nil
	}
	return value
}

func (m *Machine) PopOperator() Value {
	operator := m.operators[0]
	m.operators = m.operators[1:]
	return operator
}

func (m *Machine) step() {
	operator := m.PopOperator()
	m.evalValue(operator)
}
