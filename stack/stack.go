package stack

// Interpreter has a single method Interpret
type Interpreter interface {
	Interpret() error
}

// Stack contains cmds
type Stack struct {
	cmds []Interpreter
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
	return &Stack{}
}

func (s *Stack) Add(cmd Interpreter) {
	s.cmds = append(s.cmds, cmd)
}

func (s *Stack) Interpret() error {
	for _, cmd := range s.cmds {
		cmd.Interpret()
	}

	return nil
}
