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
		return Exception("string plus only compatible with string or int")
	}
}

func (n String) Minus(other Value) Value {
	return Exception("minus operator is not supported for string")
}

func (n String) ToString() Value {
	return n
}
