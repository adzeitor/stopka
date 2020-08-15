package stopka

type Value interface {
	// for overloading, like Integer(5) * List(1,2,3) = List(5, 10, 15)
	Plus(element Value) Value
	Minus(element Value) Value
	ToString() Value
}
