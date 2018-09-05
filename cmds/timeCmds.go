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

func (c *NowCmd) Interpret() error {
	*c.t = time.Now()
	return nil
}

// Now provides the current time
func (c *TimeCmds) Now() time.Time {
	t := new(time.Time)
	c.stack.Add(&NowCmd{t})
	return *t
}

// NewTimeCmds creates a SQLCmds
func NewTimeCmds(stack *stack.Stack) *TimeCmds {
	return &TimeCmds{stack}
}
