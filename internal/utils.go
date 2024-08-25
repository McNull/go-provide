package internal

import (
	"fmt"
	"reflect"
)

func genericTypeToString[T any]() string {
	t := reflect.TypeFor[T]() //getGenericType[T]()
	return getFullTypeName(t)
}

func getDefaultValue[T any]() T {
	var defaultValue T
	return defaultValue
}

// func getGenericType[T any]() reflect.Type {
// 	return reflect.TypeOf((*T)(nil)).Elem()
// }

func getFullTypeName(t reflect.Type) string {
	if t.PkgPath() != "" {
		return fmt.Sprintf("%s.%s", t.PkgPath(), t.Name())
	}
	return t.Name()
}
