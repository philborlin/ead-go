package log

import (
	"github.com/philborlin/ead/stack"
	log "github.com/sirupsen/logrus"
)

// InfoCmd represents an insert command
type InfoCmd struct {
	format string
	a      []interface{}
}

// Info logs at the info level
func Info(format string, a ...interface{}) {
	stack.Add(&InfoCmd{format, a})
}

// Interpret interprets the InfoCmd
func (c *InfoCmd) Interpret() error {
	log.Infof(c.format, c.a...)
	return nil
}
