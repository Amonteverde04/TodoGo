package reflection

import (
	"reflect"
	"strconv"
	"strings"
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

// Returns a slice full of property values.
func ReflectValues[T any](param T) []string {
	slice := []string{}
	t := reflect.TypeOf(param)
	v := reflect.ValueOf(param)
	for i := 0; i < t.NumField(); i++ {
		value := v.Field(i)
		if strings.Contains(value.String(), "Status") {
			slice = append(slice, strconv.Itoa(int(value.Int())))
		} else {
			slice = append(slice, v.Field(i).String())
		}
	}

	return slice
}
