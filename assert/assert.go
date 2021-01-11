package assert

import (
	"reflect"
	"testing"
)

func Equal(t *testing.T, want, got interface{}) {
	t.Helper()
	if isEmptySlices(want, got) {
		return
	}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("not equal want=%+v got=%+v", want, got)
	}
}

func isEmptySlices(first, second interface{}) bool {
	v1 := reflect.ValueOf(first)
	v2 := reflect.ValueOf(second)

	if v1.Type() != v2.Type() {
		return false
	}
	if v1.Kind() != reflect.Slice {
		return false
	}

	return (v1.Len() == 0) && (v2.Len() == 0)
}
