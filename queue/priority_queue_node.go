package queue

// PriorityQueueNode is a node in a priority queue.
type PriorityQueueNode[T any] struct {
	value    T
	priority int
}

// NewPriorityQueueNode returns a new priority queue node.
func NewPriorityQueueNode[T any](value T, priority int) PriorityQueueNode[T] {
	return PriorityQueueNode[T]{value: value, priority: priority}
}

// Value returns the value of the node.
func (node PriorityQueueNode[T]) Value() T {
	return node.value
}

// Priority returns the priority of the node.
func (node PriorityQueueNode[T]) Priority() int {
	return node.priority
}

// SetPriority sets the priority of the node.
func (node *PriorityQueueNode[T]) SetPriority(priority int) {
	node.priority = priority
}
