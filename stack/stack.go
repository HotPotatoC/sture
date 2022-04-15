package stack

import (
	"github.com/HotPotatoC/sture/linkedlist"
)

// Stack is a stack.
type Stack[T any] struct {
	list *linkedlist.LinkedList[T]
	cmp  func(T, T) int
}

// NewStack returns a new stack.
func NewStack[T any](cmp func(T, T) int) *Stack[T] {
	return &Stack[T]{
		list: linkedlist.NewLinkedList(cmp),
	}
}

// Add adds a new node to the top of the stack.
func (s *Stack[T]) Add(value T) {
	s.list.Append(value)
}

// Pop removes the top node from the stack.
func (s *Stack[T]) Pop() {
	s.list.Pop()
}

// Peek returns the value of the top node in the stack.
func (s *Stack[T]) Peek() T {
	return s.list.Tail().Value()
}

// IsEmpty returns true if the stack is empty.
func (s *Stack[T]) IsEmpty() bool {
	return s.list.IsEmpty()
}
