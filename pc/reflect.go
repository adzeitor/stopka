package pc

import (
	"reflect"
)

func interfacesToSlice(values []interface{}, kind interface{}) interface{} {
	result := reflect.MakeSlice(reflect.TypeOf(kind), 0, len(values))
	for _, value := range values {
		result = reflect.Append(result, reflect.ValueOf(value))
	}
	return result.Interface()
}
