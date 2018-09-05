package cmds

import (
	"github.com/philborlin/ead/stack"
	log "github.com/sirupsen/logrus"
)

// LoggingCmds represents a Kafka command creator
type LoggingCmds struct {
	stack *stack.Stack
}

// InfoCmd represents an insert command
type InfoCmd struct {
	format string
	a      []interface{}
}

func (c *InfoCmd) Interpret() error {
	log.Infof(c.format, c.a...)
	return nil
}

// Info logs at the info level
func (c *LoggingCmds) Info(format string, a ...interface{}) {
	c.stack.Add(&InfoCmd{format, a})
}

// NewLoggingCmds creates a LoggingCmds
func NewLoggingCmds(stack *stack.Stack) *LoggingCmds {
	return &LoggingCmds{stack}
}
