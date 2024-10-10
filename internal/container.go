package internal

import (
	"errors"
	"fmt"
)

var rootContainer = newContainer()

type Container struct {
	factories map[string]any
	instances map[string]any
	// parent    *Container
	// children  []*Container
}

func newContainer() *Container {
	return &Container{
		factories: make(map[string]any),
		instances: make(map[string]any),
		// parent:    nil,
		// children:  []*Container{},
	}
}

func (c *Container) set(key string, value any) {
	c.factories[key] = value
}

func (c *Container) get(key string) (any, error) {
	instance, ok := c.factories[key]
	if !ok {
		msg := fmt.Sprintf("factory \"%s\" not found in container", key)
		return nil, errors.New(msg)
	}
	return instance, nil
}

func (c *Container) setInstance(key string, value any) {
	c.instances[key] = value
}

func (c *Container) getInstance(key string) (any, bool) {
	instance, ok := c.instances[key]
	return instance, ok
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
