package stopka

type List []Value

func (n List) Plus(other Value) Value {
	if _, ok := other.(List); !ok {
		return append(n, other)
	}
	return append(other.(List), n...)
}

func (n List) Minus(other Value) Value {
	return Exception("minus is not defined for lists")
}

func (n List) ToString() Value {
	return Exception("conversion to string is not defined for lists")
}
