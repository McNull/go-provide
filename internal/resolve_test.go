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

	t.Run("function result should be cached", func(t *testing.T) {
		type testService struct {
			Val int
		}

		factory := func() *testService {
			return &testService{Val: 1}
		}

		tp := reflect.TypeFor[*testService]()
		tn := getFullTypeName(tp)

		container := newContainer()
		container.set(tn, factory)

		r1, err := _resolve(tp, container, invoke)

		if err != nil {
			t.Errorf("unexpected error while resolving r1: %v", err)
			return
		}

		r2, err := _resolve(tp, container, invoke)

		if err != nil {
			t.Errorf("unexpected error while resolving r2: %v", err)
			return
		}

		if r1 != r2 {
			t.Errorf("r1 and r2 should be the same")
		}
	})

	t.Run("function should return function result", func(t *testing.T) {

		// type to register the function
		type registerFuncType func() int

		want := 1

		innerFunc := func() int {
			return want
		}

		// note: the constructor function must return the type that is being registered
		constructorFunc := func() registerFuncType {
			return innerFunc
		}

		tp := reflect.TypeFor[registerFuncType]()
		tn := getFullTypeName(tp)

		container := newContainer()
		container.set(tn, constructorFunc)

		r, err := _resolve(tp, container, invoke)

		if err != nil {
			t.Errorf("unexpected error: %v", err)
			return
		}

		if r == nil {
			t.Errorf("unexpected nil result")
			return
		}

		got := r.(registerFuncType)()

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("function should return function as result twice", func(t *testing.T) {

		// register a function that returns a function

		want := 1

		innerFunc := func() int {
			return want
		}

		// type to register the function
		// function returning func() int
		type registerFuncType func() func() int

		toRegisterFunc := func() func() int {
			return innerFunc
		}

		// note: the constructor function must return the type that is being registered
		constructorFunc := func() registerFuncType {
			return toRegisterFunc
		}

		tp := reflect.TypeFor[registerFuncType]()
		tn := getFullTypeName(tp)

		container := newContainer()
		container.set(tn, constructorFunc)

		for i := 0; i < 2; i++ {
			r, err := _resolve(tp, container, invoke)

			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}

			if r == nil {
				t.Errorf("unexpected nil result")
				return
			}

			fn := r.(registerFuncType)()

			got := fn()

			if got != want {
				t.Errorf("got %d, want %d", got, want)
			}
		}
	})

}
