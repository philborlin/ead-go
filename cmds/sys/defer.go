package sys

import "github.com/philborlin/ead/stack"

// DeferCmd holds a cmd
type DeferCmd struct {
	cmd func() error
}

// Cmd holds a cmd
type Cmd struct {
	cmd func() error
}

// Defer adds a cmd to the stack
func Defer(block func() error) {
	cmd := &DeferCmd{cmd: block}
	stack.Add(cmd)
}

// Interpret adds a new command to the end of the stack
// so that this command gets deferred until the end
func (c *DeferCmd) Interpret() error {
	cmd := &Cmd{cmd: c.cmd}
	stack.Add(cmd)
	return nil
}

// Interpret actually runs the command
func (c *Cmd) Interpret() error {
	return c.cmd()
}
