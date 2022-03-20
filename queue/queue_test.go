package queue_test

import (
	"testing"

	"github.com/HotPotatoC/sture/queue"
)

func TestQueue_Enqueue(t *testing.T) {
	q := queue.NewQueue[string]()

	q.Enqueue("Adam")
	q.Enqueue("John")
	q.Enqueue("Bob")

	if q.IsEmpty() {
		t.Error("Queue should not be empty")
	}

	if q.Peek() != "Adam" {
		t.Errorf("Expected: Adam\nGot: %s", q.Peek())
	}
}

func TestQueue_Dequeue(t *testing.T) {
	q := queue.NewQueue[string]()

	q.Enqueue("Adam")
	q.Enqueue("John")
	q.Enqueue("Bob")

	q.Dequeue()

	if q.IsEmpty() {
		t.Error("Queue should not be empty")
	}

	if q.Peek() != "John" {
		t.Errorf("Expected: John\nGot: %s", q.Peek())
	}
}