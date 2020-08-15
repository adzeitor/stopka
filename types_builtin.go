package stopka

type Builtin struct {
	name string
	fn   func()
}

func (n Builtin) Plus(other Value) Value {
	panic("not implemented")
}

func (n Builtin) Minus(other Value) Value {
	panic("not implemented")
}

func (n Builtin) ToString() Value {
	return String(n.name)
}
