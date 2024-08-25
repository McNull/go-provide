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

func Get[K any]() (K, error) {
	t := reflect.TypeFor[K]()
	n := getFullTypeName(t)

	return GetByName[K](n)
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
