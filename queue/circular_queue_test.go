package queue_test

import (
	"testing"

	"github.com/HotPotatoC/sture/queue"
)

func TestCircularQueue_Enqueue(t *testing.T) {
	cq := queue.NewCircularQueue[int](6)

	ok := cq.Enqueue(1)
	if !ok {
		t.Errorf("Expected true, got false")
	}

	ok = cq.Enqueue(2)
	if !ok {
		t.Errorf("Expected true, got false")
	}

	ok = cq.Enqueue(3)
	if !ok {
		t.Errorf("Expected true, got false")
	}

	ok = cq.Enqueue(4)
	if !ok {
		t.Errorf("Expected true, got false")
	}

	ok = cq.Enqueue(5)
	if !ok {
		t.Errorf("Expected true, got false")
	}

	ok = cq.Enqueue(6)
	if ok {
		t.Errorf("Expected false, got true")
	}

	exp := "[1 2 3 4 5 0]"

	if cq.Front() != 1 {
		t.Errorf("Expected front %d, got %d", 1, cq.Front())
	}

	if cq.String() != exp {
		t.Errorf("Expected %s, got %s", exp, cq.String())
	}
}

func TestCircularQueue_Dequeue(t *testing.T) {
	cq := queue.NewCircularQueue[int](6)

	ok := cq.Enqueue(1)
	if !ok {
		t.Errorf("Expected true, got false")
	}
	ok = cq.Enqueue(2)
	if !ok {
		t.Errorf("Expected true, got false")
	}
	ok = cq.Enqueue(3)
	if !ok {
		t.Errorf("Expected true, got false")
	}
	ok = cq.Enqueue(4)
	if !ok {
		t.Errorf("Expected true, got false")
	}

	if cq.Dequeue() != 1 {
		t.Errorf("Expected %d, got %d", 1, cq.Dequeue())
	}
	if cq.Dequeue() != 2 {
		t.Errorf("Expected %d, got %d", 2, cq.Dequeue())
	}
	if cq.Dequeue() != 3 {
		t.Errorf("Expected %d, got %d", 3, cq.Dequeue())
	}
	if cq.Dequeue() != 4 {
		t.Errorf("Expected %d, got %d", 4, cq.Dequeue())
	}
}
