package cmds

import "github.com/philborlin/ead/stack"

// LoggingCmds represents a Kafka command creator
type LoggingCmds struct {
	stack *stack.Stack
}

// InfoCmd represents an insert command
type InfoCmd struct {
	format string
	a      []interface{}
}

// Info logs at the info level
func (c *LoggingCmds) Info(format string, a ...interface{}) {
	c.stack.AddCmd <- InfoCmd{format, a}
}

// NewLoggingCmds creates a LoggingCmds
func NewLoggingCmds(stack *stack.Stack) *LoggingCmds {
	return &LoggingCmds{stack}
}
