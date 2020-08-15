package stopka

type Symbol string

func (n Symbol) Plus(other Value) Value {
	if _, isList := other.(List); isList {
		return other.Plus(Identifier(n))
	}
	if otherSymbol, isSymbol := other.(Symbol); isSymbol {
		return List{Identifier(otherSymbol), Identifier(n)}
	}
	return List{other, Identifier(n)}
}

func (n Symbol) Minus(other Value) Value {
	return Exception("minus is not defined for symbols")
}

func (n Symbol) ToString() Value {
	return String(n)
}

func (n Symbol) String() string {
	return `'` + string(n)
}
