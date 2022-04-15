package stack_test

import (
	"testing"

	"github.com/HotPotatoC/sture"
	"github.com/HotPotatoC/sture/stack"
)

func BenchmarkStack_Add(b *testing.B) {
	s := stack.NewStack(sture.Compare[int])
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(i)
	}
}

func BenchmarkStack_Pop(b *testing.B) {
	s := stack.NewStack(sture.Compare[int])

	bc := []struct {
		name string
		fn   func(b *testing.B, s *stack.Stack[int])
	}{
		{name: "100k", fn: benchmarkStackPop100k},
		{name: "1m", fn: benchmarkStackPop1m},
		{name: "10m", fn: benchmarkStackPop10m},
	}
	b.ResetTimer()
	for _, bb := range bc {
		b.Run(bb.name, func(b *testing.B) {
			bb.fn(b, s)
		})
	}
}

func benchmarkStackPop100k(b *testing.B, s *stack.Stack[int]) {
	for i := 0; i < 100000; i++ {
		s.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func benchmarkStackPop1m(b *testing.B, s *stack.Stack[int]) {
	for i := 0; i < 1000000; i++ {
		s.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}

func benchmarkStackPop10m(b *testing.B, s *stack.Stack[int]) {
	for i := 0; i < 10000000; i++ {
		s.Add(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Pop()
	}
}
