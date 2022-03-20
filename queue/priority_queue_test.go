package queue_test

import (
	"testing"

	"github.com/HotPotatoC/sture/queue"
)

func TestPriorityQueue_Enqueue(t *testing.T) {
	q := queue.NewPriorityQueue[string]()
	exp := "[John]-[Adam]-[Bob]"

	q.Enqueue("Adam", 3)
	q.Enqueue("John", 2)
	q.Enqueue("Bob", 5)

	if q.IsEmpty() {
		t.Error("Queue should not be empty")
	}

	if q.Peek() != "Bob" {
		t.Errorf("Expected: Bob\nGot: %s", q.Peek())
	}

	if q.String() != exp {
		t.Errorf("\nExpected: %s\nGot: %s", exp, q.String())
	}
}

func TestPriorityQueue_Dequeue(t *testing.T) {
	q := queue.NewPriorityQueue[string]()

	q.Enqueue("Adam", 3)
	q.Enqueue("John", 2)
	q.Enqueue("Bob", 5)

	q.Dequeue()

	exp := "[John]-[Adam]"

	if q.IsEmpty() {
		t.Error("Queue should not be empty")
	}

	if q.Peek() != "Adam" {
		t.Errorf("Expected: Adam\nGot: %s", q.Peek())
	}

	if q.String() != exp {
		t.Errorf("\nExpected: %s\nGot: %s", exp, q.String())
	}
}
