package stopka

type Builtin struct {
	name string
	fn   func()
}

func (n Builtin) Plus(other Value) Value {
	return Exception("plus operator is not supported for builtin functions")
}

func (n Builtin) Minus(other Value) Value {
	return Exception("minus operator is not supported for builtin functions")
}

func (n Builtin) ToString() Value {
	return String(n.name)
}
