package cmds

// Stack contains cmds
type Stack struct {
	cmds   []interface{}
	AddCmd chan<- interface{}
}

// NewStack creates a new stack
func NewStack() *Stack {
	// Add Handlers
	return newStack()
}

// NewTestStack creates a new testing stack
func NewTestStack() *Stack {
	// Add Test Handlers
	return newStack()
}

func newStack() *Stack {
	addCmd := make(chan interface{})
	var cmds []interface{}

	go func(chan interface{}, []interface{}) {
		select {
		case cmd := <-addCmd:
			cmds = append(cmds, cmd)
		}
	}(addCmd, cmds)

	return &Stack{cmds, addCmd}
}

func (s *Stack) Interpret() error {
	return nil
}
