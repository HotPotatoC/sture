package queue

import "fmt"

// Deque is a double ended queue that supports adding and removing elements
// from both ends of the queue.
type Deque[V any] struct {
	list []V
}

// NewDeque returns a new queue.
func NewDeque[V any]() *Deque[V] {
	return &Deque[V]{
		list: make([]V, 0),
	}
}

// IsEmpty returns true if the queue is empty.
func (dq *Deque[V]) IsEmpty() bool {
	return len(dq.list) == 0
}

// Size returns the number of nodes in the queue.
func (dq *Deque[V]) Size() int {
	return len(dq.list)
}

// PushFront adds a new node to the front of the queue.
func (dq *Deque[V]) PushFront(value V) {
	dq.list = append([]V{value}, dq.list...)
}

// PushRear adds a new node to the rear of the queue.
func (dq *Deque[V]) PushRear(value V) {
	dq.list = append(dq.list, value)
}

// PopFront removes the first node from the queue.
func (dq *Deque[V]) PopFront() {
	dq.list = dq.list[1:]
}

// PopRear removes the last node from the queue.
func (dq *Deque[V]) PopRear() {
	dq.list = dq.list[:len(dq.list)-1]
}

// String returns a string representation of the queue.
func (dq *Deque[V]) String() string {
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