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

// GetByName takes a name and returns the resolved value set earlier with SetByName.
func GetByName[K any](name string) (K, error) {
	return internal.GetByName[K](name)
}

// Invoke takes a function argument, tries to resolve all argument dependencies and
// returns the result of the function.
func Invoke(fn any) (any, error) {
	return internal.Invoke(fn)
}
