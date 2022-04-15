package queue

import "fmt"

// Deque is a double ended queue that supports adding and removing elements
// from both ends of the queue.
type Deque[T any] struct {
	list []T
}

// NewDeque returns a new queue.
func NewDeque[T any]() *Deque[T] {
	return &Deque[T]{
		list: make([]T, 0),
	}
}

// IsEmpty returns true if the queue is empty.
func (dq *Deque[T]) IsEmpty() bool {
	return len(dq.list) == 0
}

// Size returns the number of nodes in the queue.
func (dq *Deque[T]) Size() int {
	return len(dq.list)
}

// PushFront adds a new node to the front of the queue.
func (dq *Deque[T]) PushFront(value T) {
	dq.list = append([]T{value}, dq.list...)
}

// PushRear adds a new node to the rear of the queue.
func (dq *Deque[T]) PushRear(value T) {
	dq.list = append(dq.list, value)
}

// PopFront removes the first node from the queue.
func (dq *Deque[T]) PopFront() {
	dq.list = dq.list[1:]
}

// PopRear removes the last node from the queue.
func (dq *Deque[T]) PopRear() {
	dq.list = dq.list[:len(dq.list)-1]
}

// String returns a string representation of the queue.
func (dq *Deque[T]) String() string {
	if dq.IsEmpty() {
		return "[]"
	}

	var s string
	for i := 0; i < len(dq.list) - 1; i++ {
		s += fmt.Sprintf("%v ", dq.list[i])
	}

	s += fmt.Sprintf("%v", dq.list[len(dq.list) - 1])

	return "[" + s + "]"
}