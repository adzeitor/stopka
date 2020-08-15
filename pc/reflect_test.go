package pc

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_interfacesToSlice(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3}, interfacesToSlice([]interface{}{1, 2, 3}, []int{}))
	assert.Equal(t, []rune{'f', 'o', 'o'}, interfacesToSlice([]interface{}{'f', 'o', 'o'}, []rune{}))

	// interface should remains interfaces
	type K interface{}
	type Foo struct{}
	type Bar struct{}
	assert.Equal(t,
		[]K{K(Foo{}), K(Foo{}), K(Bar{})},
		interfacesToSlice([]interface{}{K(Foo{}), K(Foo{}), K(Bar{})}, []K{}),
	)
}
