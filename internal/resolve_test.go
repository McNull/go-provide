package internal

import (
	"reflect"
	"testing"

	"github.com/mcnull/go-provide/testpkg"
)

func Test_resolve(t *testing.T) {

	t.Run("resolve directly", func(t *testing.T) {

		testStruct := testpkg.TestStruct{}

		container := newContainer()

		types := []struct {
			v any
		}{
			{v: 1},
			{v: "test"},
			{v: testStruct},
		}

		for _, tt := range types {
			t := reflect.TypeOf(tt.v)
			n := getFullTypeName(t)
			container.set(n, tt.v)
		}

		tests := []struct {
			name    string
			t       reflect.Type
			want    any
			wantErr bool
		}{
			{
				name:    "int",
				t:       reflect.TypeOf(1),
				want:    1,
				wantErr: false,
			},
			{
				name:    "string",
				t:       reflect.TypeOf("test"),
				want:    "test",
				wantErr: false,
			},
			{
				name:    "struct",
				t:       reflect.TypeOf(testStruct),
				want:    testStruct,
				wantErr: false,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {

				// beforeEach()

				got, err := _resolve(tt.t, container, invoke)

				if (err != nil) != tt.wantErr {
					t.Errorf("_resolve() error = %v, wantErr %v", err, tt.wantErr)
					return
				}

				// if invokeCalled != tt.wantInvoke {
				// 	t.Errorf("_resolve() invokeCalled = %v, wantInvoke %v", invokeCalled, tt.wantInvoke)
				// 	return
				// }

				if !reflect.DeepEqual(got, tt.want) {
					t.Errorf("_resolve() = %v, want %v", got, tt.want)
				}
			})
		}

	})

	t.Run("resolve from function", func(t *testing.T) {

		type testStruct struct {
			Val int
		}

		NewTesttestStruct := func() testStruct {
			return testStruct{Val: 1}
		}

		tp := reflect.TypeFor[testStruct]()
		tn := getFullTypeName(tp)

		container := newContainer()
		container.set(tn, NewTesttestStruct)

		r, err := _resolve(tp, container, invoke)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		got := r.(testStruct)

		want := NewTesttestStruct()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("_resolve() = %v, want %v", got, want)
		}

	})

	t.Run("resolve custom type", func(t *testing.T) {
		type customType string

		n := getFullTypeName(reflect.TypeFor[customType]())

		container := newContainer()
		container.set(n, customType("test"))

		r, err := _resolveByName(n, container, invoke)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		got := r.(customType)

		want := customType("test")

		if !reflect.DeepEqual(got, want) {
			t.Errorf("_resolve() = %v, want %v", got, want)
		}
	})
}
