package main

import (
	"github.com/philborlin/ead/cmds"
	"github.com/philborlin/ead/stack"
)

func main() {
	stack := stack.NewStack()
	time := cmds.NewTimeCmds(stack)
	rnd := cmds.NewRandCmds(stack)
	log := cmds.NewLoggingCmds(stack)

	i := rnd.Int()
	t := time.Now()
	log.Info("Rand: %d @ %v\n", i, t)

	stack.Interpret()
}
