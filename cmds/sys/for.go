package sys

import "github.com/philborlin/ead/stack"

type ForCmd struct {
	start     int
	condition func() bool
	increment int
	block     func(index int) error
}

func For(start int, condition func() bool, increment int, block func(index int) error) {
	cmd := &ForCmd{start, condition, increment, block}
	stack.Add(cmd)
}

func (c *ForCmd) Interpret() error {
	for i := c.start; c.condition(); i = i + c.increment {
		err := c.block(i)
		if err != nil {
			return err
		}
	}

	return nil
}

type RangeOverSliceCmd struct {
	slice []interface{}
	block func(index int, item interface{}) error
}

func RangeOverSlice(slice []interface{}, block func(index int, item interface{}) error) {
	stack.Add(&RangeOverSliceCmd{slice, block})
}

func (c *RangeOverSliceCmd) Interpret() error {
	for i, v := range c.slice {
		err := c.block(i, v)
		if err != nil {
			return err
		}
	}

	return nil
}

type RangeOverMapCmd struct {
	m     map[interface{}]interface{}
	block func(key interface{}, value interface{}) error
}

func RangeOverMap(m map[interface{}]interface{}, block func(key interface{}, value interface{}) error) {
	stack.Add(&RangeOverMapCmd{m, block})
}

func (c *RangeOverMapCmd) Interpret() error {
	for k, v := range c.m {
		err := c.block(k, v)
		if err != nil {
			return err
		}
	}

	return nil
}

type RangeOverStringCmd struct {
	s     string
	block func(index int, r rune) error
}

func RangeOverString(s string, block func(index int, r rune) error) {
	stack.Add(&RangeOverStringCmd{s, block})
}

func (c *RangeOverStringCmd) Interpret() error {
	for i, r := range c.s {
		err := c.block(i, r)
		if err != nil {
			return err
		}
	}

	return nil
}
