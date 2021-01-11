package stopka

import "fmt"

type List []Value

func (n List) Plus(other Value) Value {
	switch value := other.(type) {
	case List:
		return append(value, n...)
	case Identifier:
		return append(n, value)
	default:
		return Exception(
			fmt.Sprintf("plus is not defined for %T type", other),
		)
	}
}

func (n List) Minus(other Value) Value {
	return Exception("minus is not defined for lists")
}

func (n List) ToString() Value {
	return Exception("conversion to string is not defined for lists")
}
