package cmds

import (
	"github.com/philborlin/ead/stack"
)

// RandCmds represents a Rand command creator
type RandCmds struct {
	stack *stack.Stack
}

// IntCmd represents an insert command
type IntCmd struct {
	returnValue *int
}

// Int returns a non-negative pseudo-random int from the default Source.
func (c *RandCmds) Int() int {
	i := new(int)
	c.stack.AddCmd <- IntCmd{i}
	return *i
}

type IntnCmd struct {
	max         int
	returnValue *int
}

func (c *RandCmds) Intn(max int) int {
	i := new(int)
	c.stack.AddCmd <- IntnCmd{max, i}
	return *i
}

// NewRandCmds creates a RandCmds
func NewRandCmds(stack *stack.Stack) *RandCmds {
	return &RandCmds{stack}
}
