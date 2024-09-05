package provide

import (
	"reflect"

	"github.com/mcnull/go-provide/internal"
)

// Set takes a type argument and a value and stores the value in the container.
func Set[K any, V any](value V) {
	internal.Set[K](value)
}

// SetByName takes a name and a value and stores the value in the container.
func SetByName(name string, value any) {
	internal.SetByName(name, value)
}

// SetByType takes a type and a value and stores the value in the container.
func SetByType(t reflect.Type, value any) {
	internal.SetByType(t, value)
}

// Get takes a type argument and returns the resolved value set earlier by any of the Set* function.
func Get[K any]() (K, error) {
	return internal.Get[K]()
}

// GetByType takes a type and returns the resolved value set earlier by any of the Set* function.
func GetByType(t reflect.Type) (any, error) {
	return internal.GetByType(t)
}

// GetByName takes a name and returns the resolved value set earlier by any of the Set* function.
func GetByName[K any](name string) (K, error) {
	return internal.GetByName[K](name)
}

// GetValue resolves the generic type K and sets the resolved value to the pointer dst.
func GetValue[K any](dst *K) error {
	return internal.GetValue[K](dst)
}

// GetValueByName takes a name and a pointer to a value and resolves the value set earlier by any of the Set* function.
func GetValueByName[K any](name string, dst *K) error {
	return internal.GetValueByName(name, dst)
}

// GetValueByType takes a type and a pointer to a value and resolves the value set earlier by any of the Set* function.
func GetValueByType(t reflect.Type, dst any) error {
	return internal.GetValueByType(t, dst)
}

// GetValuesByTypes takes a slice of types and a pointer to a slice of values and resolves the values set earlier by any of the Set* function.
func GetValuesByTypes(types []reflect.Type, dst *[]any) error {
	return internal.GetValuesByTypes(types, dst)
}

// Invoke takes a function argument, tries to resolve all argument dependencies and
// returns the result of the function.
func Invoke(fn any) (any, error) {
	return internal.Invoke(fn)
}

// InvokeValue takes a function argument, tries to resolve all argument dependencies and
// sets the result of the function to the pointer dst.
func InvokeValue[T any](fn any, dst *T) error {
	return internal.InvokeValue(fn, dst)
}
