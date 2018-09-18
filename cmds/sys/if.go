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

// IfError is a special case of If that will only run if the err is != nil
// It does not accept a Else block to promote idiomatic coding practices
func IfError(err *error, block func()) {
	If(func() bool { return &err != nil }, block)
}

func If(conditional func() bool, block func()) *IfBlock {
	ifCmd := &IfCmd{}
	addIfCmd(ifCmd, conditional, block)
	stack.Add(ifCmd)
	return &IfBlock{ifCmd: ifCmd}
}

func (b *IfBlock) ElseIf(conditional func() bool, block func()) *IfBlock {
	addIfCmd(b.ifCmd, conditional, block)
	return b
}

func (b *IfBlock) Else(block func()) {
	addIfCmd(b.ifCmd, func() bool { return true }, block)
}

func addIfCmd(ifCmd *IfCmd, conditional func() bool, block func()) {
	conditionalCmd := &conditionalCmd{conditional: conditional, block: block}
	ifCmd.cmds = append(ifCmd.cmds, conditionalCmd)
}

func (c *IfCmd) Interpret() error {
	for _, cmd := range c.cmds {
		if cmd.conditional() {
			stack.NewBlock()
			cmd.block()
			err := stack.Interpret()
			if err != nil {
				return err
			}
			return nil
		}
	}

	return nil
}
