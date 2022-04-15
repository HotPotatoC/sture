package queue

import (
	"fmt"
)

// RingBuffer is a queue that supports circular insertion and removal.
type RingBuffer[T any] struct {
	list  []T
	cap   int
	front int
	rear  int
}

// NewRingBuffer creates a new circular queue with the given capacity.
func NewRingBuffer[T any](cap int) *RingBuffer[T] {
	if cap <= 0 {
		return nil
	}

	return &RingBuffer[T]{
		list:  make([]T, cap),
		cap:   cap,
		front: 0,
		rear:  0,
	}
}

// Front returns the value of the first node in the queue.
func (cq *RingBuffer[T]) Front() T {
	return cq.list[cq.front]
}

// Rear returns the value of the last node in the queue.
func (cq *RingBuffer[T]) Rear() T {
	return cq.list[cq.rear]
}

// IsFull returns true if the queue is full.
func (cq *RingBuffer[T]) IsFull() bool {
	return cq.front == (cq.rear+1)%cq.cap
}

// IsEmpty returns true if the queue is empty.
func (cq *RingBuffer[T]) IsEmpty() bool {
	return cq.rear == cq.front
}

// Enqueue adds a new node at the rear of the queue.
func (cq *RingBuffer[T]) Enqueue(value T) bool {
	if cq.IsFull() {
		return false
	}

	cq.list[cq.rear] = value
	cq.rear = (cq.rear + 1) % cq.cap

	return true
}

// Dequeue removes the first node from the queue.
func (cq *RingBuffer[T]) Dequeue() T {
	if cq.IsEmpty() {
		return *new(T) // return nil
	}

	value := cq.list[cq.front]
	cq.list[cq.front] = *new(T) // clear the value
	cq.front = (cq.front + 1) % cq.cap

	return value
}

// String returns a string representation of the queue.
func (cq *RingBuffer[T]) String() string {
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