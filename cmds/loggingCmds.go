package cmds

// LoggingCmds represents a Kafka command creator
type LoggingCmds struct {
	stack *Stack
}

// InfoCmd represents an insert command
type InfoCmd struct {
	format string
	a      []interface{}
}

// Info logs at the info level
func (c *LoggingCmds) Info(format string, a ...interface{}) int {
	i := new(int)
	c.stack.AddCmd <- InfoCmd{format, a}
	return *i
}

// NewLoggingCmds creates a LoggingCmds
func NewLoggingCmds(stack *Stack) *LoggingCmds {
	return &LoggingCmds{stack}
}
