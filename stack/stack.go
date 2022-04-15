package stack

import (
	"github.com/HotPotatoC/sture/linkedlist"
)

// Stack is a stack.
type Stack[V any] struct {
	list *linkedlist.LinkedList[V]
	cmp  func(V, V) int
}

// NewStack returns a new stack.
func NewStack[V any](cmp func(V, V) int) *Stack[V] {
	return &Stack[V]{
		list: linkedlist.NewLinkedList(cmp),
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
