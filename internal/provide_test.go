package internal

import (
	"testing"
)

type person struct {
	Name string
}

func (p *person) GetName() string {
	return p.Name
}

func (p *person) SetName(name string) {
	p.Name = name
}

type namer interface {
	GetName() string
	SetName(name string)
}

func TestSetGet(t *testing.T) {

	t.Run("struct value", func(t *testing.T) {
		myStruct := person{Name: "myStruct"}

		Set[person](myStruct)

		result, err := Get[person]()

		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if result.Name != myStruct.Name {
			t.Fatalf("Expected value %v, but got: %v", myStruct, result)
		}
	})

	t.Run("struct pointer value", func(t *testing.T) {
		myStruct := person{Name: "myStruct"}

		Set[*person](&myStruct)

		result, err := Get[*person]()

		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if result.Name != myStruct.Name {
			t.Fatalf("Expected value %v, but got: %v", myStruct, result)
		}
	})

	t.Run("interface value", func(t *testing.T) {

		myStruct := person{Name: "myStruct"}

		Set[namer](&myStruct)

		result, err := Get[namer]()

		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if result.GetName() != myStruct.GetName() {
			t.Fatalf("Expected value %v, but got: %v", myStruct, result)
		}
	})

	t.Run("custom type", func(t *testing.T) {

		type customType string

		s1 := customType("My custom type")

		Set[customType](s1)

		result, err := Get[customType]()

		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if result != "My custom type" {
			t.Fatalf("Expected value %v, but got: %v", "customType", result)
		}
	})

	t.Run("GetSetValueByName", func(t *testing.T) {

		t.Run("struct value", func(t *testing.T) {
			myStruct := person{Name: "myStruct"}

			SetByName("person", myStruct)

			var result person

			err := GetValueByName[person]("person", &result)

			if err != nil {
				t.Fatalf("Expected no error, but got: %v", err)
			}

			if result.Name != myStruct.Name {
				t.Fatalf("Expected value %v, but got: %v", myStruct, result)
			}
		})

		t.Run("struct pointer value", func(t *testing.T) {
			myStruct := person{Name: "myStruct"}

			SetByName("*person", &myStruct)

			var result *person

			err := GetValueByName("*person", &result)

			if err != nil {
				t.Fatalf("Expected no error, but got: %v", err)
			}

			if result.Name != myStruct.Name {
				t.Fatalf("Expected value %v, but got: %v", myStruct, result)
			}
		})

		t.Run("interface value", func(t *testing.T) {

			myStruct := person{Name: "myStruct"}

			SetByName("namer", &myStruct)

			var result namer

			err := GetValueByName("namer", &result)

			if err != nil {
				t.Fatalf("Expected no error, but got: %v", err)
			}

			if result.GetName() != myStruct.GetName() {
				t.Fatalf("Expected value %v, but got: %v", myStruct, result)
			}
		})

	})

	t.Run("GetValue", func(t *testing.T) {

		t.Run("struct value", func(t *testing.T) {
			myStruct := person{Name: "myStruct"}

			Set[person](myStruct)

			var result person

			err := GetValue(&result)

			if err != nil {
				t.Fatalf("Expected no error, but got: %v", err)
			}

			if result.Name != myStruct.Name {
				t.Fatalf("Expected value %v, but got: %v", myStruct, result)
			}
		})

		t.Run("struct pointer value", func(t *testing.T) {
			myStruct := person{Name: "myStruct"}

			Set[*person](&myStruct)

			var result *person

			err := GetValue(&result)

			if err != nil {
				t.Fatalf("Expected no error, but got: %v", err)
			}

			if result.Name != myStruct.Name {
				t.Fatalf("Expected value %v, but got: %v", myStruct, result)
			}
		})

		t.Run("interface value", func(t *testing.T) {

			myStruct := person{Name: "myStruct"}

			Set[namer](&myStruct)

			var result namer

			err := GetValue(&result)

			if err != nil {
				t.Fatalf("Expected no error, but got: %v", err)
			}

			if result.GetName() != myStruct.GetName() {
				t.Fatalf("Expected value %v, but got: %v", myStruct, result)
			}
		})
	})
}
