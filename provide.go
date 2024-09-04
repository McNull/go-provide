package provide

import (
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

// Get takes a type argument and returns the resolved value set earlier with Set.
func Get[K any]() (K, error) {
	return internal.Get[K]()
}

// GetValue takes a pointer to a value and returns the resolved value set earlier with Set.
func GetValue[K any](value *K) error {
	return internal.GetValue[K](value)
}

// GetByName takes a name and returns the resolved value set earlier with Set or SetByName.
func GetByName[K any](name string) (K, error) {
	return internal.GetByName[K](name)
}

// GetValueByName takes a name and a pointer to a value and returns the resolved value set earlier with Set or SetByName.
func GetValueByName[K any](name string, value *K) error {
	return internal.GetValueByName(name, value)
}

// Invoke takes a function argument, tries to resolve all argument dependencies and
// returns the result of the function.
func Invoke(fn any) (any, error) {
	return internal.Invoke(fn)
}
