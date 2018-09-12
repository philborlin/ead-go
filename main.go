package main

import (
	"github.com/philborlin/ead/cmds"
	"github.com/philborlin/ead/stack"
	"github.com/philborlin/ead/system"
)

func main() {
	mainStack := stack.NewStack()
	time := cmds.NewTimeCmds(mainStack)
	rnd := cmds.NewRandCmds(mainStack)
	sys := system.NewSystemCmds(mainStack)

	i := rnd.Int()
	t := time.Now()

	sys.Eif(func() bool { return t.Second()%2 == 0 }, func() *stack.Stack {
		ifStack := stack.NewStack()
		cmds.NewLoggingCmds(ifStack).Info("Even Rand: %d @ %v\n", i, t)
		return ifStack
	}()).Eelse(func() *stack.Stack {
		ifStack := stack.NewStack()
		cmds.NewLoggingCmds(ifStack).Info("Odd Rand: %d @ %v\n", i, t)
		return ifStack
	}())

	mainStack.Interpret()
}
