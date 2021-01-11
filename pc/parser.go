package pc

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"unicode"
)

type State struct {
	Value   interface{}
	Remains string
	Err     error
}

// FIXME: input is state???
type Parser func(state State) State

var (
	ErrNoMatch       = errors.New("no match")
	ErrUnexpectedEnd = errors.New("unexpected end of line")
)

func Fail(err error) Parser {
	return func(s State) State {
		s.Err = err
		return s
	}
}

func (parser Parser) Run(str string) State {
	state := parser(State{Remains: str})
	return state
}

func (parser Parser) Map(fn interface{}) Parser {
	return func(s State) State {
		state := parser(s)
		if state.Err != nil {
			return state
		}
		result := reflect.ValueOf(fn).Call([]reflect.Value{reflect.ValueOf(state.Value)})
		state.Value = result[0].Interface()
		return state
	}
}

func OneOf(parsers ...Parser) Parser {
	return func(s State) State {
		for _, parser := range parsers {
			state := parser(s)
			if state.Err != nil {
				continue
			}
			return state
		}
		return Fail(ErrNoMatch)(s)
	}
}

func Satisfy(fn func(rune) bool) Parser {
	return func(s State) State {
		if len(s.Remains) == 0 {
			return State{Err: ErrUnexpectedEnd}
		}
		r := []rune(s.Remains)[0]
		if fn(r) {
			// FIXME: multibyte bug
			return State{Value: r, Remains: string([]rune(s.Remains)[1:])}
		}
		return State{Remains: s.Remains, Err: ErrNoMatch}
	}
}

func Char(expect string) Parser {
	return Satisfy(func(r rune) bool {
		return strings.ContainsRune(expect, r)
	})
}

func NotChar(notExpect string) Parser {
	return Satisfy(func(r rune) bool {
		return !strings.ContainsRune(notExpect, r)
	})
}

func Str(expect string) Parser {
	return func(s State) State {
		if strings.HasPrefix(s.Remains, expect) {
			return State{Value: expect, Remains: s.Remains[len(expect):]}
		}
		return State{Remains: s.Remains, Err: ErrNoMatch}
	}
}

func Integer() Parser {
	parseInt := func(value []rune) int {
		num, _ := strconv.Atoi(string(value))
		return num
	}
	return Many1(Digit(), []rune{}).Map(parseInt)
}

func Digit() Parser {
	return func(s State) State {
		r := []rune(s.Remains)[0]
		if unicode.Is(unicode.Digit, r) {
			return State{Value: r, Remains: s.Remains[1:]}
		}
		return State{Err: errors.New("not digit")}
	}
}

func Many1(parser Parser, kind interface{}) Parser {
	return func(s State) (state State) {
		state = State{Remains: s.Remains}
		var (
			values []interface{}
			count  int
		)

		for len(s.Remains) > 0 {
			nextState := parser(s)
			if nextState.Err != nil {
				break
			}
			s.Remains = nextState.Remains
			values = append(values, nextState.Value)
			count++
		}
		if count == 0 {
			return State{Remains: state.Remains, Err: ErrNoMatch}
		}
		state.Remains = s.Remains
		if values != nil {
			state.Value = interfacesToSlice(values, kind)
		}
		return state
	}
}

func Many(parser Parser, kind interface{}) Parser {
	return func(s State) (state State) {
		state = State{Remains: s.Remains}
		var values []interface{}

		for len(s.Remains) > 0 {
			nextState := parser(s)
			if nextState.Err != nil {
				break
			}
			s.Remains = nextState.Remains
			values = append(values, nextState.Value)
		}
		state.Remains = s.Remains
		state.Value = interfacesToSlice(values, kind)
		return state
	}
}

// FIXME: rename to Success
func Pure(value interface{}) Parser {
	return func(state State) State {
		state.Value = value
		return state
	}
}

func Between(left Parser, parser Parser, right Parser) Parser {
	return left.DiscardLeft(parser.DiscardRight(right))
}

func (parser Parser) Between(left Parser, right Parser) Parser {
	return Between(left, parser, right)
}

// FIXME: use lift?
func (parser Parser) DiscardLeft(right Parser) Parser {
	return func(state State) State {
		leftState := parser(state)
		if leftState.Err != nil {
			return leftState
		}
		return right(leftState)
	}
}

// FIXME: use lift?
func (parser Parser) DiscardRight(right Parser) Parser {
	return func(state State) State {
		leftState := parser(state)
		if leftState.Err != nil {
			return leftState
		}
		rightState := right(leftState)
		rightState.Value = leftState.Value
		return rightState
	}
}

func runesToString(value interface{}) interface{} {
	return string(value.([]rune))
}

func NotQuotedWord() Parser {
	return Many1(NotChar(` "'`), []rune{}).
		Map(runesToString)
}
