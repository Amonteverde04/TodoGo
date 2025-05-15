package reflection

import (
	"reflect"
)

// Returns a slice full of property names.
func ReflectProperties[T any](param T) []string {
	slice := []string{}
	t := reflect.TypeOf(param)
	for i := 0; i < t.NumField(); i++ {
		slice = append(slice, t.Field((i)).Name)
	}

	return slice
}

// Returns a slice full of property names.
func ReflectValues[T any](param T) []string {
	slice := []string{}
	t := reflect.TypeOf(param)
	v := reflect.ValueOf(param)
	for i := 0; i < t.NumField(); i++ {
		slice = append(slice, v.Field(i).String())
	}

	return slice
}
