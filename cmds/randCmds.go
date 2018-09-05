package cmds

import (
	"math/rand"

	"github.com/philborlin/ead/stack"
)

// RandCmds represents a Rand command creator
type RandCmds struct {
	stack *stack.Stack
}

// IntCmd represents an integer command
type IntCmd struct {
	returnValue *int
}

func (c *IntCmd) Interpret() error {
	*c.returnValue = rand.Int()
	return nil
}

// Int returns a non-negative pseudo-random int from the default Source.
func (c *RandCmds) Int() *int {
	i := new(int)
	c.stack.Add(&IntCmd{returnValue: i})
	return i
}

type IntnCmd struct {
	max         int
	returnValue *int
}

func (c *RandCmds) Intn(max int) *int {
	i := new(int)
	c.stack.Add(&IntnCmd{max: max, returnValue: i})
	return i
}

func (c *IntnCmd) Interpret() error {
	*c.returnValue = rand.Intn(c.max)
	return nil
}

// NewRandCmds creates a RandCmds
func NewRandCmds(stack *stack.Stack) *RandCmds {
	return &RandCmds{stack}
}
