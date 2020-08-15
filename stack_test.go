package stopka

import "testing"

func TestStack(t *testing.T) {
	stack := Stack{}
	stack.Push(Integer(5))
	stack.Push(Integer(3))

	top, _ := stack.Pop()
	if top != Integer(3) {
		t.Error("top element must be 3, got", top)
	}

	top, _ = stack.Pop()
	if top != Integer(5) {
		t.Error("after pop next element must be 5, got", top)
	}
}
