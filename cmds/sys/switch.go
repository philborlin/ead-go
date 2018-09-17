package sys

import "github.com/philborlin/ead/stack"

// TODO Add fallthrough

type switchConditional = func(value interface{}) bool

func ValueConditional(value interface{}) switchConditional {
	return func(actual interface{}) bool { return actual == value }
}

type switchConditionalCmd struct {
	conditional switchConditional
	block       func()
}

type SwitchCmd struct {
	cmds  []*switchConditionalCmd
	value interface{}
}

type SwitchBlock struct {
	switchCmd *SwitchCmd
}

func Switch(value interface{}) *SwitchBlock {
	cmd := &SwitchCmd{value: value}
	stack.Add(cmd)
	return &SwitchBlock{switchCmd: cmd}
}

func (c *SwitchCmd) Interpret() error {
	for _, cmd := range c.cmds {
		if cmd.conditional(c.value) {
			stack.NewBlock()
			cmd.block()
			err := stack.Interpret()
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (b *SwitchBlock) Case(conditionals []switchConditional, block func()) *SwitchBlock {
	addSwitchCmd(b.switchCmd, conditionals, block)
	return b
}

func (b *SwitchBlock) Default(block func()) {
	addSwitchCmd(b.switchCmd, []switchConditional{func(v interface{}) bool { return true }}, block)
}

func addSwitchCmd(cmd *SwitchCmd, conditionals []switchConditional, block func()) {
	for _, conditional := range conditionals {
		conditionalCmd := &switchConditionalCmd{conditional: conditional, block: block}
		cmd.cmds = append(cmd.cmds, conditionalCmd)
	}
}
