package stack_test

import (
	"testing"

	"github.com/HotPotatoC/sture"
	"github.com/HotPotatoC/sture/stack"
)

func TestAdd(t *testing.T) {
	s := stack.NewStack(sture.Compare[int])
	s.Add(1)
	s.Add(2)
	s.Add(3)

	if s.Peek() != 3 {
		t.Errorf("Expected 3, got %v", s.Peek())
	}

	if s.IsEmpty() {
		t.Errorf("Expected stack to not be empty")
	}
}

func TestPop(t *testing.T) {
	s := stack.NewStack(sture.Compare[int])
	s.Add(1)
	s.Add(2)
	s.Add(3)

	s.Pop()

	if s.Peek() != 2 {
		t.Errorf("Expected 2, got %v", s.Peek())
	}

	if s.IsEmpty() {
		t.Errorf("Expected stack to not be empty")
	}

	s.Pop()

	if s.Peek() != 1 {
		t.Errorf("Expected 2, got %v", s.Peek())
	}

	if s.IsEmpty() {
		t.Errorf("Expected stack to not be empty")
	}

	s.Pop()

	if !s.IsEmpty() {
		t.Errorf("Expected stack to be empty")
	}
}
