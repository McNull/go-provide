package internal

import (
	"reflect"
)

// resolve resolves a type from the provided container.
// if the type is a function, it will be invoked and the result will be returned.
func resolve(t reflect.Type, container *Container) (any, error) {
	return _resolve(t, container, invoke)
}

// Used for testing resolve function
func _resolve(t reflect.Type, container *Container, invoke invokeFunc) (any, error) {
	n := getFullTypeName(t)
	return _resolveByName(n, container, invoke)
}

func _resolveByName(name string, container *Container, invoke invokeFunc) (any, error) {
	v, err := container.get(name)

	if err != nil {
		return nil, err
	}

	// if the value is a function, invoke it
	vt := reflect.TypeOf(v)
	if vt.Kind() == reflect.Func {
		return invoke(v, container)
	}

	return v, nil
}

type resolveAllFunc func([]reflect.Type, *Container) ([]reflect.Value, error)

// resolveAll resolves all types in a slice from the provided container.
func resolveAll(tt []reflect.Type, container *Container) ([]reflect.Value, error) {
	vv := make([]reflect.Value, len(tt))

	for i, t := range tt {
		v, err := resolve(t, container)
		if err != nil {
			return nil, err
		}

		vv[i] = reflect.ValueOf(v)
	}

	return vv, nil
}
