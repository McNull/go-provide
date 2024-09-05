package internal

import (
	"reflect"
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

type count struct {
	count int
}

func (c *count) Increment() {
	c.count++
}

func (c *count) GetCount() int {
	return c.count
}

type counter interface {
	Increment()
	GetCount() int
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

func TestSetByType(t *testing.T) {
	t.Run("struct value", func(t *testing.T) {
		value := person{Name: "myStruct"}
		tp := reflect.TypeOf(value)

		SetByType(tp, value)

		result, err := Get[person]()

		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if result.Name != value.Name {
			t.Fatalf("Expected value %v, but got: %v", value, result)
		}
	})

	t.Run("struct pointer value", func(t *testing.T) {
		value := person{Name: "myStruct"}
		tp := reflect.TypeOf(&value)

		SetByType(tp, &value)

		result, err := Get[*person]()

		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if result.Name != value.Name {
			t.Fatalf("Expected value %v, but got: %v", value, result)
		}
	})

	t.Run("interface value", func(t *testing.T) {

		value := person{Name: "myStruct"}
		tp := reflect.TypeFor[namer]()

		SetByType(tp, &value)

		result, err := Get[namer]()

		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if result.GetName() != value.GetName() {
			t.Fatalf("Expected value %v, but got: %v", value, result)
		}
	})
}

func TestGetByType(t *testing.T) {

	t.Run("struct value", func(t *testing.T) {
		value := person{Name: "myStruct"}
		tp := reflect.TypeOf(value)

		Set[person](value)

		result, err := GetByType(tp)

		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if result.(person).Name != value.Name {
			t.Fatalf("Expected value %v, but got: %v", value, result)
		}
	})

	t.Run("struct pointer value", func(t *testing.T) {
		value := person{Name: "myStruct"}
		tp := reflect.TypeOf(&value)

		Set[*person](&value)

		result, err := GetByType(tp)

		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if result.(*person).Name != value.Name {
			t.Fatalf("Expected value %v, but got: %v", value, result)
		}
	})

	t.Run("interface value", func(t *testing.T) {

		value := person{Name: "myStruct"}
		tp := reflect.TypeFor[namer]()

		Set[namer](&value)

		result, err := GetByType(tp)

		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if result.(namer).GetName() != value.GetName() {
			t.Fatalf("Expected value %v, but got: %v", value, result)
		}
	})
}

func TestGetValueByType(t *testing.T) {

	t.Run("struct value", func(t *testing.T) {

		value := person{Name: "myStruct"}
		tp := reflect.TypeOf(value)

		Set[person](value)

		var result person

		err := GetValueByType(tp, &result)

		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if result.Name != value.Name {
			t.Fatalf("Expected value %v, but got: %v", value, result)
		}
	})

	t.Run("struct pointer value", func(t *testing.T) {

		value := person{Name: "myStruct"}
		tp := reflect.TypeOf(&value)

		Set[*person](&value)

		var result *person

		err := GetValueByType(tp, &result)

		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if result.Name != value.Name {
			t.Fatalf("Expected value %v, but got: %v", value, result)
		}
	})

	t.Run("interface value", func(t *testing.T) {

		value := person{Name: "myStruct"}
		tp := reflect.TypeFor[namer]()

		Set[namer](&value)

		var result namer

		err := GetValueByType(tp, &result)

		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if result.GetName() != value.GetName() {
			t.Fatalf("Expected value %v, but got: %v", value, result)
		}
	})

}

func TestGetValuesByTypes(t *testing.T) {

	t.Run("struct values", func(t *testing.T) {

		p1 := person{Name: "myStruct"}
		tp1 := reflect.TypeOf(p1)

		Set[person](p1)

		p2 := count{count: 10}
		tp2 := reflect.TypeOf(p2)

		Set[count](p2)

		types := []reflect.Type{tp1, tp2}
		dst := []any{}

		err := GetValuesByTypes(types, &dst)

		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if len(dst) != 2 {
			t.Fatalf("Expected length 2, but got: %v", len(dst))
		}

		{
			result := dst[0].(person)

			if result.Name != p1.Name {
				t.Fatalf("Expected value %v, but got: %v", p1, result)
			}
		}

		{
			result := dst[1].(count)

			if result.count != p2.count {
				t.Fatalf("Expected value %v, but got: %v", p2, result)
			}
		}
	})

	t.Run("struct pointer values", func(t *testing.T) {

		p1 := person{Name: "myStruct"}
		tp1 := reflect.TypeOf(&p1)

		Set[*person](&p1)

		p2 := count{count: 10}
		tp2 := reflect.TypeOf(&p2)

		Set[*count](&p2)

		types := []reflect.Type{tp1, tp2}
		dst := []any{}

		err := GetValuesByTypes(types, &dst)

		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if len(dst) != 2 {
			t.Fatalf("Expected length 2, but got: %v", len(dst))
		}

		{
			result := dst[0].(*person)

			if result.Name != p1.Name {
				t.Fatalf("Expected value %v, but got: %v", p1, result)
			}
		}

		{
			result := dst[1].(*count)

			if result.count != p2.count {
				t.Fatalf("Expected value %v, but got: %v", p2, result)
			}
		}
	})

	t.Run("interface values", func(t *testing.T) {

		p1 := person{Name: "myStruct"}
		tp1 := reflect.TypeFor[namer]()

		Set[namer](&p1)

		p2 := count{count: 10}
		tp2 := reflect.TypeFor[counter]()

		Set[counter](&p2)

		types := []reflect.Type{tp1, tp2}
		dst := []any{}

		err := GetValuesByTypes(types, &dst)

		if err != nil {
			t.Fatalf("Expected no error, but got: %v", err)
		}

		if len(dst) != 2 {
			t.Fatalf("Expected length 2, but got: %v", len(dst))
		}

		{
			result := dst[0].(namer)

			if result.GetName() != p1.GetName() {
				t.Fatalf("Expected value %v, but got: %v", p1, result)
			}
		}

		{
			result := dst[1].(counter)

			if result.GetCount() != p2.GetCount() {
				t.Fatalf("Expected value %v, but got: %v", p2, result)
			}
		}
	})
}
