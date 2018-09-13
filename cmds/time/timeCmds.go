package time

import (
	"time"

	"github.com/philborlin/ead/stack"
)

// NowCmd represents an now command
type NowCmd struct {
	t *time.Time
}

// Now provides the current time
func Now() time.Time {
	t := new(time.Time)
	stack.Add(&NowCmd{t})
	return *t
}

// Interpret inteprets the cmd
func (c *NowCmd) Interpret() error {
	*c.t = time.Now()
	return nil
}
