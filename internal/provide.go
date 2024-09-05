package internal

import (
	"fmt"
	"reflect"
)

func Set[K any, V any](value V) {
	k := genericTypeToString[K]()
	SetByName(k, value)
}

func SetByName(name string, value any) {
	rootContainer.set(name, value)
}

func SetByType(t reflect.Type, value any) {
	name := getFullTypeName(t)
	SetByName(name, value)
}

func Get[K any]() (K, error) {
	t := reflect.TypeFor[K]()
	n := getFullTypeName(t)

	return GetByName[K](n)
}

func GetByType(t reflect.Type) (any, error) {
	n := getFullTypeName(t)
	return GetByName[any](n)
}

func GetByName[K any](name string) (K, error) {
	instance, err := _resolveByName(name, rootContainer, invoke)

	if err != nil {
		return getDefaultValue[K](), err
	}

	s, ok := instance.(K)

	if ok {
		return s, nil
	}

	return getDefaultValue[K](), fmt.Errorf("type assertion from %T to %T failed", instance, genericTypeToString[K]())
}

func GetValue[K any](dst *K) error {
	instance, err := Get[K]()

	if err != nil {
		return err
	}

	*dst = instance

	return nil
}

func GetValueByName[K any](name string, dst *K) error {
	instance, err := GetByName[K](name)

	if err != nil {
		return err
	}

	*dst = instance

	return nil
}

func GetValueByType(t reflect.Type, dst any) error {
	// dst should always be a pointer

	if dst == nil {
		return fmt.Errorf("dst is nil")
	}

	if !isPointer(dst) {
		return fmt.Errorf("dst is not a pointer")
	}

	instance, err := GetByType(t)

	if err != nil {
		return err
	}

	// assign the instance to dst

	reflect.ValueOf(dst).Elem().Set(reflect.ValueOf(instance))

	return nil
}

func GetValuesByTypes(types []reflect.Type, dst *[]any) error {

	if dst == nil {
		return fmt.Errorf("dst is nil")
	}

	if len(types) != len(*dst) {
		*dst = make([]any, len(types))
	}

	for i, t := range types {
		err := GetValueByType(t, &(*dst)[i])

		if err != nil {
			return fmt.Errorf("error getting value for type \"%s\": %v", t, err)
		}
	}

	return nil
}
