package main

import (
	"github.com/philborlin/ead/cmds"
	"github.com/prometheus/common/log"
)

func main() {
	stack := stack()

	kafka.Append(nil)
	id := sql.Insert("INSERT INFO foo (bar) VALUES (?);", "qux")
	log.Info("Last id: %d", id)

	stack.Interpret()
}

func stack() *Stack {
	stack := cmds.NewStack()
	kafka := cmds.NewKafkaCmds(stack)
	log := cmds.NewLoggingCmds(stack)
	sql := cmds.NewSQLCmds(stack)
}
