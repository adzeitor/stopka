package stopka

import (
	"testing"

	"github.com/adzeitor/stopka/assert"
)

func assertStack(t *testing.T, m *Machine, elements ...Value) {
	assert.Equal(t, elements, m.Stack())
}

func assertOperators(t *testing.T, m *Machine, operators ...Value) {
	assert.Equal(t, operators, m.Operators())
}

func TestMachine_Step(t *testing.T) {
	m := New().LoadOperators("1 5 swap dup drop +")
	assertOperators(t, m, Integer(1), Integer(5), Identifier("swap"), Identifier("dup"), Identifier("drop"), Identifier("+"))

	// 1
	m.step()
	assertStack(t, m, Integer(1))
	assertOperators(t, m, Integer(5), Identifier("swap"), Identifier("dup"), Identifier("drop"), Identifier("+"))

	// 5
	m.step()
	assertStack(t, m, Integer(1), Integer(5))
	assertOperators(t, m, Identifier("swap"), Identifier("dup"), Identifier("drop"), Identifier("+"))

	// swap
	m.step()
	assertStack(t, m, Integer(5), Integer(1))
	assertOperators(t, m, Identifier("dup"), Identifier("drop"), Identifier("+"))

	// dup
	m.step()
	assertStack(t, m, Integer(5), Integer(1), Integer(1))
	assertOperators(t, m, Identifier("drop"), Identifier("+"))

	// drop
	m.step()
	assertStack(t, m, Integer(5), Integer(1))
	assertOperators(t, m, Identifier("+"))

	// +
	m.step()
	assertStack(t, m, Integer(6))
}

func TestMachine_Eval(t *testing.T) {
	tests := []struct {
		name   string
		line   string
		want   []Value
		halted bool
	}{
		{
			line: `40 2 +`,
			want: Values{Integer(42)},
		},
		{
			line: `84 42 -`,
			want: Values{Integer(42)},
		},
		{
			line: `44 2 - 100 +`,
			want: Values{Integer(142)},
		},
		{
			line: `84 42 '- apply`,
			want: Values{Integer(42)},
		},
		{
			line: `1 5 swap dup drop +`,
			want: Values{Integer(6)},
		},
		{
			name: "string concat",
			line: `"foo" "bar" +`,
			want: Values{
				String("foobar"),
			},
		},
		{
			name: "list concat",
			line: `(1 2) (4 5) +`,
			want: Values{
				List{Integer(1), Integer(2), Integer(4), Integer(5)},
			},
		},
		//{
		//	name: "map with quoted word",
		//	line: `(1 2 3) 'string map`,
		//	want: Values{
		//		List{String("1"), String("2"), String("3")},
		//	},
		//},
		{
			name: "map with quoted list",
			line: `(1 2 3) (dup +) map`,
			want: Values{
				List{Integer(2), Integer(4), Integer(6)},
			},
		},
		{
			name: "symbols",
			line: `'+ '- 'dup 'foo`,
			want: Values{
				Symbol("+"), Symbol("-"), Symbol("dup"), Symbol("foo"),
			},
		},
		{
			name: "plus on symbols creates list",
			line: `'dup '* + 'swap +`,
			want: Values{
				List{
					Identifier("dup"),
					Identifier("*"),
					Identifier("swap"),
				},
			},
		},
		{
			name:   "empty stack plus",
			line:   `+ - swap dup`,
			halted: true,
		},
		{
			name:   "unknown identifier",
			line:   `42 24 super_new_identifier`,
			want:   Values{Integer(42), Integer(24)},
			halted: true,
		},
		{
			name:   "list is not support - operator",
			line:   `(1 2 3) (4 5 6) -`,
			want:   Values{},
			halted: true,
		},
		// FIXME:
		//{
		//	name:   "apply with non-function argument",
		//	line:   `(1 2 3) 5 apply`,
		//	want:   Values{},
		//	halted: true,
		//},
		{
			// FIXME: is it should be error?
			name: "map with non-function argument",
			line: `(1 2 3) (5) map`,
			want: Values{
				List{Integer(5), Integer(5), Integer(5)},
			},
			halted: false,
		},
		{
			name: "define simple value",
			line: `42 'answer define answer answer`,
			want: Values{
				Integer(42), Integer(42),
			},
			halted: false,
		},
		{
			name: "map with defined function",
			line: `(dup +) 'double define (1 2 3) double map`,
			want: Values{
				List{Integer(2), Integer(4), Integer(6)},
			},
		},
	}
	for _, tt := range tests {
		name := tt.name
		if name == "" {
			name = tt.line
		}
		t.Run(name, func(t *testing.T) {
			machine := New().Eval(tt.line)
			assert.Equal(t, tt.want, machine.Stack())
			assert.Equal(t, tt.halted, machine.IsHalted())
		})
	}
}

func TestMachine_EvalUnsupportedOverload(t *testing.T) {
	unsupportedExpressions := []string{
		`(1 2 3) string`,
		`5 (1 2 3) -`,
		`"foo" "bar" -`,
		`"foo" 42 +`,
		`"foo" 42 -`,
		`(1 2 3) 5 -`,
		`(1 2 3) 5 +`,
		`5 (1 2 3) +`,
		`(1 2 3) (4 5 6) -`,
		`'foo 5 -`,
	}
	for _, tt := range unsupportedExpressions {
		t.Run(tt, func(t *testing.T) {
			machine := New().Eval(tt)
			assert.Equal(t, true, machine.IsHalted())
		})
	}
}
