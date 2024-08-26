package internal

import (
	"reflect"
	"testing"

	"github.com/mcnull/go-provide/testpkg"
)

func TestGetFullTypeName(t *testing.T) {
	tests := []struct {
		name string
		t    reflect.Type
		want string
	}{
		{
			name: "Type with package path",
			t:    reflect.TypeOf((*testing.T)(nil)).Elem(),
			want: "testing.T",
		},
		{
			name: "Type without package path",
			t:    reflect.TypeOf(0),
			want: "int",
		},
		{
			name: "Type from another package",
			t:    reflect.TypeOf((*testpkg.Person)(nil)).Elem(),
			want: "github.com/mcnull/go-provide/testpkg.Person",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFullTypeName(tt.t); got != tt.want {
				t.Errorf("getFullTypeName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetFullTypeName2(t *testing.T) {

	t.Run("External type", func(t *testing.T) {
		tt := reflect.TypeFor[testpkg.Person]()
		want := "github.com/mcnull/go-provide/testpkg.Person"
		got := getFullTypeName(tt)

		if got != want {
			t.Errorf("getFullTypeName() = \"%v\", want \"%v\"", got, want)
		}
	})

	t.Run("External type with pointer", func(t *testing.T) {
		tt := reflect.TypeFor[*testpkg.Person]()
		want := "*github.com/mcnull/go-provide/testpkg.Person"
		got := getFullTypeName(tt)

		if got != want {
			t.Errorf("getFullTypeName() = \"%v\", want \"%v\"", got, want)
		}
	})

	t.Run("External type with pointer to pointer", func(t *testing.T) {
		tt := reflect.TypeFor[**testpkg.Person]()
		want := "**github.com/mcnull/go-provide/testpkg.Person"
		got := getFullTypeName(tt)

		if got != want {
			t.Errorf("getFullTypeName() = \"%v\", want \"%v\"", got, want)
		}
	})

}
