package internal

import (
	"errors"
	"fmt"
)

var rootContainer = newContainer()

type Container struct {
	instances map[string]any
	// parent    *Container
	// children  []*Container
}

func newContainer() *Container {
	return &Container{
		instances: make(map[string]any),
		// parent:    nil,
		// children:  []*Container{},
	}
}

func (c *Container) set(key string, value any) {
	c.instances[key] = value
}

func (c *Container) get(key string) (any, error) {
	instance, ok := c.instances[key]
	if !ok {
		// if c.parent != nil {
		// 	return c.parent.get(key)
		// }

		msg := fmt.Sprintf("instance \"%s\" not found in container", key)
		return nil, errors.New(msg)
	}
	return instance, nil
}

// func (c *Container) removeChild(child *Container) {
// 	if child == nil {
// 		return
// 	}

// 	if child.parent != c {
// 		return
// 	}

// 	for i, ch := range c.children {
// 		if ch == child {
// 			c.children = append(c.children[:i], c.children[i+1:]...)
// 			break
// 		}
// 	}

// 	child.parent = nil
// }

// func (c *Container) addChild(child *Container) {
// 	if child == nil {
// 		return
// 	}

// 	if child.parent == c {
// 		return
// 	}

// 	child.parent = c
// 	c.children = append(c.children, child)
// }
