package queue

import (
	"golang.org/x/exp/constraints"
)

// Queue is a queue.
type Queue[V constraints.Ordered] struct {
	list  []V
	nSize int
}

// NewQueue returns a new queue.
func NewQueue[V constraints.Ordered]() *Queue[V] {
	return &Queue[V]{
		list: make([]V, 0),
	}
}

// Enqueue adds a new node to the end of the queue.
func (q *Queue[V]) Enqueue(value V) {
	q.list = append(q.list, value)
	q.nSize++
}

// Dequeue removes the first node from the queue.
func (q *Queue[V]) Dequeue() {
	if q.nSize == 0 {
		return
	}

	q.list = q.list[1:]
	q.nSize--
}

// Peek returns the value of the first node in the queue.
func (q *Queue[V]) Peek() V {
	return q.list[0]
}

// IsEmpty returns true if the queue is empty.
func (q *Queue[V]) IsEmpty() bool {
	return q.nSize == 0
}
