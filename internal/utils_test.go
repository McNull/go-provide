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
