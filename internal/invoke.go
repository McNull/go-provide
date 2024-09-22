package internal

import (
	"fmt"
	"reflect"
)

// Invoke takes a function argument, tries to resolve all argument dependencies and
// returns the result of the function.
func Invoke(fn any) (any, error) {
	return invoke(fn, rootContainer)
}

func InvokeValue[T any](fn any, dst *T) error {
	v, err := invoke(fn, rootContainer)
	if err != nil {
		return err
	}
	*dst = v.(T)
	return nil
}

type invokeFunc func(fn any, container *Container) (any, error)

// invoke invokes a function with the provided container.
// The function arguments will be resolved from the container and passed to the function.
// The result of the function will be returned.
func invoke(fn any, container *Container) (any, error) {
	return _invoke(fn, container, resolveAll)
}

// used for testing.
func _invoke(fn any, container *Container, resolveAll resolveAllFunc) (any, error) {
	tt, err := getFuncArgTypes(fn)

	if err != nil {
		return nil, err
	}

	vv, err := resolveAll(tt, container)

	if err != nil {
		return nil, err
	}

	rr := reflect.ValueOf(fn).Call(vv)

	if len(rr) == 0 {
		return nil, nil
	}

	// if the function returns a single value
	if len(rr) == 1 {
		// if the value is an error, change return value to nil, error
		if rr[0].Type().Implements(reflect.TypeOf((*error)(nil)).Elem()) {
			if rr[0].IsNil() {
				return nil, nil
			}

			return nil, rr[0].Interface().(error)
		}

		return rr[0].Interface(), nil
	}

	// if the function retured 2 values, check if the second value is an error
	if len(rr) == 2 {
		if rr[1].IsNil() {
			return rr[0].Interface(), nil
		}

		return nil, rr[1].Interface().(error)
	}

	return nil, fmt.Errorf("invoking function \"%s\" returned too many values", reflect.TypeOf(fn).Name())
}

// getFuncArgTypes returns the types of the arguments of a function.
func getFuncArgTypes(fn any) ([]reflect.Type, error) {
	fnType := reflect.TypeOf(fn)

	if fnType.Kind() != reflect.Func {
		return nil, fmt.Errorf("provided value is not a function")
	}

	c := fnType.NumIn()

	tt := make([]reflect.Type, c)

	for i := 0; i < c; i++ {
		tt[i] = fnType.In(i)
	}

	return tt, nil
}
