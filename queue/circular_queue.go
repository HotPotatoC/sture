package queue

import (
	"fmt"
)

// CircularQueue is a queue that supports circular insertion and removal.
type CircularQueue[V any] struct {
	list  []V
	cap   int
	front int
	rear  int
}

// NewCircularQueue creates a new circular queue with the given capacity.
func NewCircularQueue[V any](cap int) *CircularQueue[V] {
	if cap <= 0 {
		return nil
	}

	return &CircularQueue[V]{
		list:  make([]V, cap),
		cap:   cap,
		front: 0,
		rear:  0,
	}
}

// Front returns the value of the first node in the queue.
func (cq *CircularQueue[V]) Front() V {
	return cq.list[cq.front]
}

// Rear returns the value of the last node in the queue.
func (cq *CircularQueue[V]) Rear() V {
	return cq.list[cq.rear]
}

// IsFull returns true if the queue is full.
func (cq *CircularQueue[V]) IsFull() bool {
	return cq.front == (cq.rear+1)%cq.cap
}

// IsEmpty returns true if the queue is empty.
func (cq *CircularQueue[V]) IsEmpty() bool {
	return cq.rear == cq.front
}

// Enqueue adds a new node at the rear of the queue.
func (cq *CircularQueue[V]) Enqueue(value V) bool {
	if cq.IsFull() {
		return false
	}

	cq.list[cq.rear] = value
	cq.rear = (cq.rear + 1) % cq.cap

	return true
}

// Dequeue removes the first node from the queue.
func (cq *CircularQueue[V]) Dequeue() V {
	if cq.IsEmpty() {
		return *new(V) // return nil
	}

	value := cq.list[cq.front]
	cq.list[cq.front] = *new(V) // clear the value
	cq.front = (cq.front + 1) % cq.cap

	return value
}

// String returns a string representation of the queue.
func (cq *CircularQueue[V]) String() string {
	if cq.IsEmpty() {
		return "[]"
	}

	var s string
	for i := cq.front; i != cq.rear; i = (i + 1) % cq.cap {
		s += fmt.Sprintf("%v ", cq.list[i])
	}
	s += fmt.Sprintf("%v", cq.list[cq.rear])

	return "[" + s + "]"
}