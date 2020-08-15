package stopka

type Exception string

func (e Exception) Plus(other Value) Value {
	return e
}

func (e Exception) Minus(other Value) Value {
	return e
}

func (e Exception) ToString() Value {
	return String(e)
}
