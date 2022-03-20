package queue

import (
	"golang.org/x/exp/constraints"
)

// PriorityQueueNode is a node in a priority queue.
type PriorityQueueNode[V constraints.Ordered] struct {
	value    V
	priority int
}

// NewPriorityQueueNode returns a new priority queue node.
func NewPriorityQueueNode[V constraints.Ordered](value V, priority int) PriorityQueueNode[V] {
	return PriorityQueueNode[V]{value: value, priority: priority}
}

// Value returns the value of the node.
func (node PriorityQueueNode[V]) Value() V {
	return node.value
}

// Priority returns the priority of the node.
func (node PriorityQueueNode[V]) Priority() int {
	return node.priority
}

// SetPriority sets the priority of the node.
func (node *PriorityQueueNode[V]) SetPriority(priority int) {
	node.priority = priority
}
