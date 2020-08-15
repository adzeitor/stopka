package pc

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"strconv"
	"testing"
)

func Test_parse_OneOf(t *testing.T) {
	parser := OneOf(
		Str("foo"),
		Str("bar"),
		Integer(),
	)

	state := parser.Run("foo xxx")
	assert.Equal(t, State{
		Value:   "foo",
		Remains: " xxx",
	}, state)

	state = parser.Run("bar yyy")
	assert.Equal(t, State{
		Value:   "bar",
		Remains: " yyy",
	}, state)

	state = parser.Run("55 yyy")
	assert.Equal(t, State{
		Value:   55,
		Remains: " yyy",
	}, state)

	state = parser.Run("qux")
	assert.Equal(t, State{Remains: "qux", Err: ErrNoMatch}, state)
}

func TestChar(t *testing.T) {
	tests := []struct {
		name   string
		parser Parser
		str    string
		want   State
	}{
		{
			name:   "success",
			parser: Char("h"),
			str:    "hello",
			want:   State{Value: 'h', Remains: "ello"},
		},
		{
			name:   "multi byte",
			parser: Char("п"),
			str:    "привет",
			want:   State{Value: 'п', Remains: "ривет"},
		},
		{
			name:   "empty string",
			parser: Char("п"),
			str:    "",
			want:   State{Remains: "", Err: ErrUnexpectedEnd},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := tt.parser.Run(tt.str)
			assert.Equal(t, tt.want, state)
		})
	}
}

func TestDigit(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want State
	}{
		{
			name: "digit",
			str:  "555",
			want: State{Value: '5', Remains: "55"},
		},
		{
			name: "number with suffix",
			str:  "5xyz",
			want: State{Value: '5', Remains: "xyz"},
		},
		{
			name: "not digit",
			str:  "xxx",
			want: State{Err: errors.New("not digit")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := Digit().Run(tt.str)
			assert.Equal(t, tt.want, state)
		})
	}
}

func TestMany(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want State
	}{
		{
			name: "digits",
			str:  "543",
			want: State{Value: []rune{'5', '4', '3'}, Remains: ""},
		},
		{
			name: "digits with suffix",
			str:  "54xyz",
			want: State{Value: []rune{'5', '4'}, Remains: "xyz"},
		},
		{
			name: "not digit",
			str:  "xxx",
			want: State{Remains: "xxx"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := Many(Digit(), []rune{}).Run(tt.str)
			assert.Equal(t, tt.want, state)
		})
	}
}

func TestMany1(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want State
	}{
		{
			name: "digits",
			str:  "543",
			want: State{Value: []rune{'5', '4', '3'}, Remains: ""},
		},
		{
			name: "digits with suffix",
			str:  "54xyz",
			want: State{Value: []rune{'5', '4'}, Remains: "xyz"},
		},
		{
			name: "not digit",
			str:  "xxx",
			want: State{Remains: "xxx", Err: ErrNoMatch},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state := Many1(Digit(), []rune{}).Run(tt.str)
			assert.Equal(t, tt.want, state)
		})
	}
}

func Test_parse_NotQuotedWord(t *testing.T) {
	p1 := NotQuotedWord()

	result := p1.Run(`foo 555`)
	assert.Equal(t, State{Value: "foo", Remains: " 555"}, result)
}

func TestParser_Between(t *testing.T) {
	openParen := Char("(")
	closeParen := Char(")")
	foo := Str("foo")

	parser := foo.Between(openParen, closeParen)

	assert.Equal(t, State{Value: "foo"}, parser.Run("(foo)"))
	assert.Equal(t, ErrNoMatch, parser.Run("(bar)").Err)
	assert.Equal(t, ErrUnexpectedEnd, parser.Run("(foo").Err)
	assert.Equal(t, ErrNoMatch, parser.Run("foo)").Err)
	assert.Equal(t, ErrNoMatch, parser.Run("foo").Err)
}

func TestParser_Map(t *testing.T) {
	parser := Pure(5).Map(func(n int) string {
		return strconv.Itoa(n)
	})
	assert.Equal(t, State{Value: "5", Remains: "..."}, parser.Run("..."))

	t.Run("type of value on error should be set", func(t *testing.T) {
		t.Skip()
		parser = Fail(errors.New("error")).Map(func(n int) string {
			return strconv.Itoa(n)
		})
		got := parser.Run("...")
		assert.IsType(t, "", got.Value)
	})
}