package main

import (
	"github.com/philborlin/ead/cmds/log"
	"github.com/philborlin/ead/cmds/rand"
	"github.com/philborlin/ead/cmds/sys"
	"github.com/philborlin/ead/cmds/time"
	"github.com/philborlin/ead/stack"
)

func main() {
	i := rand.Int()
	t := time.Now()

	sys.If(func() bool { return t.Second()%2 == 0 }, func() {
		log.Info("Even Rand: %d @ %v\n", i, t)
	}).Else(func() {
		log.Info("Odd Rand: %d @ %v\n", i, t)
	})

	stack.Interpret()
}
