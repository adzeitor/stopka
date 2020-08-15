package stopka

type Identifier string

func (n Identifier) Plus(other Value) Value {
	return Exception("plus is not defined for identifiers")
}

func (n Identifier) Minus(other Value) Value {
	return Exception("minus is not defined for identifiers")
}

func (n Identifier) ToString() Value {
	return String(n)
}
