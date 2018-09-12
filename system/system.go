package system

import "github.com/philborlin/ead/stack"

type SystemCmds struct {
	stack *stack.Stack
}

// NewSystemCmds creates a SystemCmds
func NewSystemCmds(stack *stack.Stack) *SystemCmds {
	return &SystemCmds{stack}
}
