package rand

import (
	"math/rand"

	"github.com/philborlin/ead/stack"
)

// IntCmd represents an integer command
type IntCmd struct {
	returnValue *int
}

// Int returns a non-negative pseudo-random int from the default Source.
func Int() *int {
	i := new(int)
	stack.Add(&IntCmd{returnValue: i})
	return i
}

func (c *IntCmd) Interpret() error {
	*c.returnValue = rand.Int()
	return nil
}

type IntnCmd struct {
	max         int
	returnValue *int
}

func Intn(max int) *int {
	i := new(int)
	stack.Add(&IntnCmd{max: max, returnValue: i})
	return i
}

func (c *IntnCmd) Interpret() error {
	*c.returnValue = rand.Intn(c.max)
	return nil
}
