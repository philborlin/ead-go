package stack

import (
	"sync"

	"github.com/tylerb/gls"
)

// We should experiment with a Block() method which will create a new Interpreter list
// This will work in a stack like function where Block() will push() a new list
// and Interpret() will execute the top list and then pop() the list off
// This will allow us to setup one stack per file
// * cmd setup will look more like imports so it will look more natural
// * we can make conditionals and loops look much more natural

// Interpreter has a single method Interpret
type Interpreter interface {
	Interpret() error
}

type block struct {
	cmds []Interpreter
}

func newBlock() *block {
	return &block{}
}

type effectStack struct {
	stack *blockStack
}

func newEffectStack() *effectStack {
	return &effectStack{stack: newBlockStack()}
}

// Get gives you the goroutine-local stack.
//
// WARNING: Caching this in a variable and accessing on a differnt goroutine WILL cause
// massive and unpredictable problems
func get() *effectStack {
	s := gls.Get("stack")
	if s == nil {
		gls.Set("stack", newEffectStack())
		s = gls.Get("stack")
	}
	return s.(*effectStack)
}

// Add adds a new cmd to the current block of the goroutine-local stack
func Add(cmd Interpreter) {
	s := get()
	b := s.stack.peek()
	b.cmds = append(b.cmds, cmd)
}

// Interpret interprets the current block of the goroutine-local stack
// This will pop() the block off the stack
func Interpret() error {
	s := get()
	b := s.stack.pop()

	var err error
	for _, cmd := range b.cmds {
		err = cmd.Interpret()
		if err != nil {
			break
		}
	}

	s.stack.ensureNotEmpty()
	return err
}

// NewBlock will push() a new block onto the stack. Add() and Interpret() will
// operate against the current block
func NewBlock() {
	s := get()
	s.stack.push(newBlock())
}

// --- The code below is based off of https://flaviocopes.com/golang-data-structure-stack/

type blockStack struct {
	blocks []*block
	lock   sync.RWMutex
}

func newBlockStack() *blockStack {
	s := &blockStack{}
	s.ensureNotEmpty()
	return s
}

// Push adds a block to the top of the stack
func (s *blockStack) push(t *block) {
	s.lock.Lock()
	s.blocks = append(s.blocks, t)
	s.lock.Unlock()
}

// Pop removes a block from the top of the stack and returns it
func (s *blockStack) pop() *block {
	s.lock.Lock()
	block := s.blocks[len(s.blocks)-1]
	s.blocks = s.blocks[0 : len(s.blocks)-1]
	s.lock.Unlock()
	return block
}

// Peek returns a block from the top of the stack without removing it
func (s *blockStack) peek() *block {
	s.lock.Lock()
	block := s.blocks[len(s.blocks)-1]
	s.lock.Unlock()
	return block
}

func (s *blockStack) ensureNotEmpty() {
	if len(s.blocks) == 0 {
		s.push(newBlock())
	}
}
