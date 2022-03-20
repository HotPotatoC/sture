package queue

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// PriorityQueue is a queue that supports priority-based insertion.
type PriorityQueue[V constraints.Ordered] struct {
	list  []PriorityQueueNode[V]
	nSize int
}

// NewPriorityQueue creates a new priority queue.
func NewPriorityQueue[V constraints.Ordered]() *PriorityQueue[V] {
	return &PriorityQueue[V]{
		list: make([]PriorityQueueNode[V], 0),
	}
}

// Enqueue adds a new node at a certain position in the queue based on the priority.
func (q *PriorityQueue[V]) Enqueue(value V, priority int) {
	newNode := NewPriorityQueueNode(value, priority)

	pos := q.search(priority) // search for the position to insert the new node

	q.list = append(q.list, PriorityQueueNode[V]{}) // add empty node

	copy(q.list[pos+1:], q.list[pos:]) // shift elements to the right

	q.list[pos] = newNode // insert new node at pos
	q.nSize++
}

// Dequeue removes the last node from the queue.
func (q *PriorityQueue[V]) Dequeue() {
	if q.nSize == 0 {
		return
	}

	q.list = q.list[:q.nSize-1]
	q.nSize--
}

// Peek returns the value of the node with the highest priority in the queue.
func (q *PriorityQueue[V]) Peek() V {
	return q.list[q.nSize-1].Value()
}

// IsEmpty returns true if the queue is empty.
func (q *PriorityQueue[V]) IsEmpty() bool {
	return q.nSize == 0
}

// String returns a string representation of the queue.
func (q *PriorityQueue[V]) String() string {
	var s string

	for idx, node := range q.list {
		s += fmt.Sprintf("[%v]", node.value)
		if idx != q.nSize-1 {
			s += "-"
		}
	}

	return s
}

// search returns the position of the node that is the closest to the given priority.
func (q *PriorityQueue[V]) search(priority int) int {
	low, high := 0, q.nSize
	for low < high {
		mid := int(uint(low+high) >> 1) // avoid overflow when computing mid

		if q.list[mid].Priority() <= priority {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return low
}
