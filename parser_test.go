package stopka

import (
	"testing"

	"github.com/adzeitor/stopka/assert"
	"github.com/adzeitor/stopka/pc"
)

func Test_parse(t *testing.T) {
	t.Run("integers and string", func(t *testing.T) {
		got, _ := parse(`55 66 "foo"`)
		assert.Equal(t, []Value{
			Integer(55),
			Integer(66),
			String("foo"),
		}, got)
	})

	t.Run("identifiers", func(t *testing.T) {
		got, _ := parse(`1 5 swap dup drop +`)
		assert.Equal(t, []Value{
			Integer(1),
			Integer(5),
			Identifier("swap"),
			Identifier("dup"),
			Identifier("drop"),
			Identifier("+"),
		}, got)
	})

	t.Run("symbols", func(t *testing.T) {
		got, _ := parse(`1 5 'swap 'dup 'drop '+`)
		assert.Equal(t, []Value{
			Integer(1),
			Integer(5),
			Symbol("swap"),
			Symbol("dup"),
			Symbol("drop"),
			Symbol("+"),
		}, got)
	})

	t.Run("empty list", func(t *testing.T) {
		got, _ := parse(`()`)
		assert.Equal(t, []Value{
			List{},
		}, got)
	})

	t.Run("empty lists", func(t *testing.T) {
		got, _ := parse(`1 () 2 () 3`)
		assert.Equal(t, []Value{
			Integer(1), List{}, Integer(2), List{}, Integer(3),
		}, got)
	})

	t.Run("list of one element", func(t *testing.T) {
		got, _ := parse(`(1)`)
		assert.Equal(t, []Value{
			List{Integer(1)},
		}, got)
	})

	t.Run("list", func(t *testing.T) {
		got, _ := parse(`(1 2 3)`)
		assert.Equal(t, []Value{
			List{Integer(1), Integer(2), Integer(3)},
		}, got)
	})

	t.Run("nested lists", func(t *testing.T) {
		got, _ := parse(`(1 (2 3 4) 5 (6 7) 8)`)
		assert.Equal(t, []Value{
			List{
				Integer(1),
				List{Integer(2), Integer(3), Integer(4)},
				Integer(5),
				List{Integer(6), Integer(7)},
				Integer(8),
			},
		}, got)
	})

	t.Run("more than one space", func(t *testing.T) {
		got, _ := parse(`     (   1   2   (   3  4     )    + )    `)
		assert.Equal(t, []Value{
			List{
				Integer(1), Integer(2),
				List{Integer(3), Integer(4)},
				Identifier("+"),
			},
		}, got)
	})

	t.Run("not quoted string", func(t *testing.T) {
		_, err := parse(`"foo`)
		assert.Equal(t, pc.ErrNoMatch, err)
	})

	t.Run("not closed parens of list", func(t *testing.T) {
		_, err := parse(`(1 2 3`)
		assert.Equal(t, pc.ErrNoMatch, err)
	})

	t.Run("unknown expression", func(t *testing.T) {
		_, err := parse(`[5]`)
		assert.Equal(t, pc.ErrNoMatch, err)
	})

	t.Run("empty input is not panic", func(t *testing.T) {
		_, err := parse(``)
		assert.Equal(t, pc.ErrNoMatch, err)
	})
}
