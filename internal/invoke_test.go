package internal

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetFuncArgTypes(t *testing.T) {
	tests := []struct {
		name      string
		input     any
		wantTypes []reflect.Type
		wantErr   bool
	}{
		{
			name:      "Valid function with multiple arguments",
			input:     func(a int, b string) {},
			wantTypes: []reflect.Type{reflect.TypeOf(0), reflect.TypeOf("")},
			wantErr:   false,
		},
		{
			name:      "Valid function with no arguments",
			input:     func() {},
			wantTypes: []reflect.Type{},
			wantErr:   false,
		},
		{
			name:      "Invalid input (not a function)",
			input:     123,
			wantTypes: nil,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTypes, err := getFuncArgTypes(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("getFuncArgTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotTypes, tt.wantTypes) {
				t.Errorf("getFuncArgTypes() = %v, want %v", gotTypes, tt.wantTypes)
			}
		})
	}
}

func Test_invoke(t *testing.T) {

	container := newContainer()

	type testData struct {
		name          string
		fn            any
		want          any
		wantErr       bool
		resolveResult []reflect.Value
		resolveError  error
	}

	testError := fmt.Errorf("test error")

	tests := []testData{
		{
			name:    "Invalid function",
			fn:      123,
			wantErr: true,
		},
		{
			name: "Function with no arguments",
			fn:   func() int { return 123 },
			want: 123,
		},
		{
			name: "Function returning error",
			fn: func() (any, error) {
				return nil, testError
			},
			want:    testError,
			wantErr: true,
		},
	}

	var currentTest *testData = nil

	resolveAll := func(tt []reflect.Type, container *Container) ([]reflect.Value, error) {
		return currentTest.resolveResult, currentTest.resolveError
	}

	touch := func(x any) {}

	for i, tt := range tests {

		currentTest = &tt
		t.Run(tt.name, func(t *testing.T) {

			touch(i) // golang suck

			result, err := _invoke(tt.fn, container, resolveAll)

			if tt.wantErr {
				if err == nil {
					t.Errorf("Expected error")
					return
				}
			} else {
				if err != nil {
					t.Error(err)
					return
				}

				if !reflect.DeepEqual(tt.want, result) {
					t.Errorf("Expected %v, got %v", tt.want, result)
					return
				}
			}

		})
	}

}

func TestInvokeValue(t *testing.T) {
	type MyResult struct {
		Value int
	}

	t.Run("Struct value", func(t *testing.T) {
		testFunc := func() MyResult {
			return MyResult{Value: 123}
		}

		var result MyResult

		err := InvokeValue(testFunc, &result)

		if err != nil {
			t.Error(err)
			return
		}

		if result.Value != 123 {
			t.Errorf("Expected 123, got %d", result.Value)
		}
	})

	t.Run("Struct pointer value", func(t *testing.T) {
		testFunc := func() *MyResult {
			return &MyResult{Value: 123}
		}

		var result *MyResult

		err := InvokeValue(testFunc, &result)

		if err != nil {
			t.Error(err)
			return
		}

		if result.Value != 123 {
			t.Errorf("Expected 123, got %d", result.Value)
		}
	})
}
