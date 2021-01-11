package stopka

import (
	"fmt"
	"strconv"
)

type Integer int

func (n Integer) Plus(other Value) Value {
	switch value := other.(type) {
	case Integer:
		return value + n
	default:
		return Exception(
			fmt.Sprintf("plus is not defined for %T type", other),
		)
	}
}

func (n Integer) Minus(other Value) Value {
	switch value := other.(type) {
	case Integer:
		return value - n
	default:
		return Exception(
			fmt.Sprintf("minus is not defined for %T type", other),
		)
	}
}

func (n Integer) ToString() Value {
	return String(strconv.Itoa(int(n)))
}
