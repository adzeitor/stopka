package stopka

import (
	"strconv"
)

type String string

func (n String) Plus(other Value) Value {
	switch value := other.(type) {
	case String:
		return value + n
	case Integer:
		return String(strconv.Itoa(int(value)) + string(n))
	default:
		return Exception("string plus only compatible with string or int ")
	}
}

func (n String) Minus(other Value) Value {
	panic("not implemented")
}

func (n String) ToString() Value {
	return n
}
