package sys

import (
	"github.com/philborlin/ead/stack"
)

type conditionalCmd struct {
	conditional func() bool
	block       func()
}

type IfCmd struct {
	cmds []*conditionalCmd
}

type IfBlock struct {
	ifCmd *IfCmd
}

func If(conditional func() bool, block func()) *IfBlock {
	ifCmd := &IfCmd{}
	addCmd(ifCmd, conditional, block)
	stack.Add(ifCmd)
	return &IfBlock{ifCmd: ifCmd}
}

func (b *IfBlock) ElseIf(conditional func() bool, block func()) *IfBlock {
	addCmd(b.ifCmd, conditional, block)
	return b
}

func (b *IfBlock) Else(block func()) {
	addCmd(b.ifCmd, func() bool { return true }, block)
}

func addCmd(ifCmd *IfCmd, conditional func() bool, block func()) {
	conditionalCmd := &conditionalCmd{conditional: conditional, block: block}
	ifCmd.cmds = append(ifCmd.cmds, conditionalCmd)
}

func (c *IfCmd) Interpret() error {
	var err error

	for _, cmd := range c.cmds {
		if cmd.conditional() {
			stack.NewBlock()
			cmd.block()
			err = stack.Interpret()
			break
		}
	}

	return err
}
