package internal

import (
	"testing"
)

func TestContainer(t *testing.T) {
	t.Run("set & get", func(t *testing.T) {
		container := newContainer()

		// Test case 1: Key exists in the container
		key := "myKey"
		value := "myValue"
		container.set(key, value)

		result, err := container.get(key)
		if err != nil {
			t.Errorf("Expected no error, but got: %v", err)
		}

		if result != value {
			t.Errorf("Expected value %v, but got: %v", value, result)
		}

		// Test case 2: Key does not exist in the container
		nonExistentKey := "nonExistentKey"

		_, err = container.get(nonExistentKey)
		if err == nil {
			t.Errorf("Expected error, but got nil")
		}
	})
	/*
	   t.Run("add & remove child", func(t *testing.T) {

	   		// Test case 1: Add child
	   		parent := NewContainer()
	   		child := NewContainer()

	   		parent.addChild(child)

	   		if len(parent.children) != 1 {
	   			t.Errorf("Expected 1 child, but got %d", len(parent.children))
	   		}

	   		if child.parent != parent {
	   			t.Errorf("Expected child parent to be %v, but got %v", parent, child.parent)
	   		}

	   		// Test case 2: Remove child
	   		parent.removeChild(child)

	   		if len(parent.children) != 0 {
	   			t.Errorf("Expected 0 child, but got %d", len(parent.children))
	   		}

	   		if child.parent != nil {
	   			t.Errorf("Expected child parent to be nil, but got %v", child)
	   		}

	   		// Test case 3: Remove nil child
	   		parent.removeChild(nil)

	   		if len(parent.children) != 0 {
	   			t.Errorf("Expected 0 child, but got %d", len(parent.children))
	   		}
	   	})

	   	t.Run("get should fallback to parent", func(t *testing.T) {
	   		parent := NewContainer()
	   		child := NewContainer()

	   		parent.addChild(child)

	   		key := "myKey"
	   		value := "myValue"
	   		parent.set(key, value)

	   		result, err := child.get(key)
	   		if err != nil {
	   			t.Errorf("Expected no error, but got: %v", err)
	   		}

	   		if result != value {
	   			t.Errorf("Expected value %v, but got: %v", value, result)
	   		}
	   	})
	*/
}
