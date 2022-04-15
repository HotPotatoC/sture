package queue

// Queue is a queue.
type Queue[T any] struct {
	list  []T
	nSize int
}

// NewQueue returns a new queue.
func NewQueue[T any]() *Queue[T] {
	return &Queue[T]{
		list: make([]T, 0),
	}
}

// Enqueue adds a new node to the end of the queue.
func (q *Queue[T]) Enqueue(value T) {
	q.list = append(q.list, value)
	q.nSize++
}

// Dequeue removes the first node from the queue.
func (q *Queue[T]) Dequeue() {
	if q.nSize == 0 {
		return
	}

	q.list = q.list[1:]
	q.nSize--
}

// Peek returns the value of the first node in the queue.
func (q *Queue[T]) Peek() T {
	return q.list[0]
}

// IsEmpty returns true if the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return q.nSize == 0
}
