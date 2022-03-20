package queue_test

import (
	"testing"

	"github.com/HotPotatoC/sture/queue"
)

func TestDeque_PushFront(t *testing.T) {
	d := queue.NewDeque[int]()
	exp := "[5 4 3 2 1]"

	d.PushFront(1)
	d.PushFront(2)
	d.PushFront(3)
	d.PushFront(4)
	d.PushFront(5)

	if d.IsEmpty() {
		t.Errorf("Expected not empty, got empty")
	}

	if d.Size() != 5 {
		t.Errorf("Expected size %d, got %d", 5, d.Size())
	}

	if d.String() != exp {
		t.Errorf("\nExpected: %s\nGot: %s", exp, d.String())
	}
}

func TestDeque_PushRear(t *testing.T) {
	d := queue.NewDeque[int]()
	exp := "[1 2 3 4 5]"

	d.PushRear(1)
	d.PushRear(2)
	d.PushRear(3)
	d.PushRear(4)
	d.PushRear(5)

	if d.IsEmpty() {
		t.Errorf("Expected not empty, got empty")
	}

	if d.Size() != 5 {
		t.Errorf("Expected size %d, got %d", 5, d.Size())
	}

	if d.String() != exp {
		t.Errorf("\nExpected: %s\nGot: %s", exp, d.String())
	}
}

func TestDeque_PopFront(t *testing.T) {
	d := queue.NewDeque[int]()
	exp := "[3 4 5]"

	d.PushRear(1)
	d.PushRear(2)
	d.PushRear(3)
	d.PushRear(4)
	d.PushRear(5)

	d.PopFront()
	d.PopFront()

	if d.IsEmpty() {
		t.Errorf("Expected not empty, got empty")
	}

	if d.Size() != 3 {
		t.Errorf("Expected size %d, got %d", 3, d.Size())
	}

	if d.String() != exp {
		t.Errorf("\nExpected: %s\nGot: %s", exp, d.String())
	}
}

func TestDeque_PopRear(t *testing.T) {
	d := queue.NewDeque[int]()
	exp := "[1 2 3]"

	d.PushRear(1)
	d.PushRear(2)
	d.PushRear(3)
	d.PushRear(4)
	d.PushRear(5)

	d.PopRear()
	d.PopRear()

	if d.IsEmpty() {
		t.Errorf("Expected not empty, got empty")
	}

	if d.Size() != 3 {
		t.Errorf("Expected size %d, got %d", 3, d.Size())
	}

	if d.String() != exp {
		t.Errorf("\nExpected: %s\nGot: %s", exp, d.String())
	}
}
