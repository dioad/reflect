package reflect

import (
	"reflect"
	"strings"
)

// HasField returns true it T has a field called fieldName
func HasField[T any](fieldName string) bool {
	a := *new(T)
	t := reflect.TypeOf(a)

	if t.Kind() == reflect.Struct {
		_, found := t.FieldByName(fieldName)
		return found
	}

	if t.Kind() == reflect.Pointer {
		_, found := t.Elem().FieldByName(fieldName)
		return found
	}

	return false
}

// Name returns the name of the type, including package and whether it's a pointer or not
func Name[T any]() (string, bool) {
	a := *new(T)
	t := reflect.TypeOf(a)
	name := t.String()

	isPointer := false
	if strings.HasPrefix(name, "*") {
		isPointer = true
		name = strings.Replace(name, "*", "", 1)
	}

	return name, isPointer
}
