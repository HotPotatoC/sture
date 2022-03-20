package stack

import (
	"github.com/HotPotatoC/sture/linkedlist"
	"golang.org/x/exp/constraints"
)

// Stack is a stack.
type Stack[V constraints.Ordered] struct {
	list *linkedlist.LinkedList[V]
}

// NewStack returns a new stack.
func NewStack[V constraints.Ordered]() *Stack[V] {
	return &Stack[V]{
		list: linkedlist.NewLinkedList[V](),
	}
}

// Add adds a new node to the top of the stack.
func (s *Stack[V]) Add(value V) {
	s.list.Append(value)
}

// Pop removes the top node from the stack.
func (s *Stack[V]) Pop() {
	s.list.Pop()
}

// Peek returns the value of the top node in the stack.
func (s *Stack[V]) Peek() V {
	return s.list.Tail().Value()
}

// IsEmpty returns true if the stack is empty.
func (s *Stack[V]) IsEmpty() bool {
	return s.list.IsEmpty()
}
