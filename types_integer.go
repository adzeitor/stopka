package stopka

import "strconv"

type Integer int

func (n Integer) Plus(other Value) Value {
	switch value := other.(type) {
	case Integer:
		return value + n
	default:
		panic("unknown type")
	}
}

func (n Integer) Minus(other Value) Value {
	switch value := other.(type) {
	case Integer:
		return value - n
	default:
		panic("unknown type")
	}
}

func (n Integer) ToString() Value {
	return String(strconv.Itoa(int(n)))
}
