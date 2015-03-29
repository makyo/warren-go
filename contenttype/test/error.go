package test

import (
	"fmt"
)

type Error struct{}

func (c *Error) Safe() bool {
	return false
}

func (c *Error) RenderDisplayContent(content interface{}) (string, error) {
	return "", fmt.Errorf("Error")
}

func (c *Error) RenderIndexContent(content interface{}) (string, error) {
	return "", fmt.Errorf("Error")
}
