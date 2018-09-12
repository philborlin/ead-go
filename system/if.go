package system

import (
	"github.com/philborlin/ead/stack"
)

type ConditionalCmd struct {
	conditional func() bool
	cmd         stack.Interpreter
}

type IfCmd struct {
	cmds []*ConditionalCmd
}

type IfBlock struct {
	ifCmd *IfCmd
}

func (c *SystemCmds) Eif(conditional func() bool, cmd stack.Interpreter) *IfBlock {
	ifCmd := &IfCmd{}
	addCmd(ifCmd, conditional, cmd)
	c.stack.Add(ifCmd)
	return &IfBlock{ifCmd: ifCmd}
}

func (b *IfBlock) EelseIf(conditional func() bool, cmd stack.Interpreter) *IfBlock {
	addCmd(b.ifCmd, conditional, cmd)
	return b
}

func (b *IfBlock) Eelse(cmd stack.Interpreter) {
	addCmd(b.ifCmd, func() bool { return true }, cmd)
}

func addCmd(ifCmd *IfCmd, conditional func() bool, cmd stack.Interpreter) {
	conditionalCmd := &ConditionalCmd{conditional: conditional, cmd: cmd}
	ifCmd.cmds = append(ifCmd.cmds, conditionalCmd)
}

func (c *IfCmd) Interpret() error {
	var err error

	for _, cmd := range c.cmds {
		if cmd.conditional() {
			err = cmd.cmd.Interpret()
			break
		}
	}

	return err
}
