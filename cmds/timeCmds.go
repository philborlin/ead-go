package cmds

import (
	"time"

	"github.com/philborlin/ead/stack"
)

// TimeCmds represents a Kafka command creator
type TimeCmds struct {
	stack *stack.Stack
}

// NowCmd represents an insert command
type NowCmd struct {
	t *time.Time
}

// Now provides the current time
func (c *TimeCmds) Now() time.Time {
	t := new(time.Time)
	c.stack.AddCmd <- NowCmd{t}
	return *t
}

// NewTimeCmds creates a SQLCmds
func NewTimeCmds(stack *stack.Stack) *TimeCmds {
	return &TimeCmds{stack}
}
