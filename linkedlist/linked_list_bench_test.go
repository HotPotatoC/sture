package linkedlist_test

import (
	"testing"

	"github.com/HotPotatoC/sture/linkedlist"
)

func BenchmarkLinkedList_Append(b *testing.B) {
	ll := linkedlist.NewLinkedList[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ll.Append(i)
	}
}

func BenchmarkLinkedList_PushHead(b *testing.B) {
	ll := linkedlist.NewLinkedList[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ll.PushHead(i)
	}
}

func BenchmarkLinkedList_InsertAt(b *testing.B) {
	ll := linkedlist.NewLinkedList[int]()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ll.InsertAt(i, i+1)
	}
}

func BenchmarkLinkedList_Pop(b *testing.B) {
	ll := linkedlist.NewLinkedList[int]()

	bc := []struct {
		name string
		fn   func(b *testing.B, ll *linkedlist.LinkedList[int])
	}{
		{name: "benchmarkLinkedListPop100k", fn: benchmarkLinkedListPop100k},
		{name: "benchmarkLinkedListPop1m", fn: benchmarkLinkedListPop1m},
		{name: "benchmarkLinkedListPop10m", fn: benchmarkLinkedListPop10m},
	}
	b.ResetTimer()
	for _, bb := range bc {
		b.Run(bb.name, func(b *testing.B) {
			bb.fn(b, ll)
		})
	}
}

func benchmarkLinkedListPop100k(b *testing.B, ll *linkedlist.LinkedList[int]) {
	for i := 0; i < 100000; i++ {
		ll.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ll.Pop()
	}
}

func benchmarkLinkedListPop1m(b *testing.B, ll *linkedlist.LinkedList[int]) {
	for i := 0; i < 1000000; i++ {
		ll.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ll.Pop()
	}
}

func benchmarkLinkedListPop10m(b *testing.B, ll *linkedlist.LinkedList[int]) {
	for i := 0; i < 10000000; i++ {
		ll.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ll.Pop()
	}
}

func BenchmarkLinkedList_PopHead(b *testing.B) {
	ll := linkedlist.NewLinkedList[int]()

	bc := []struct {
		name string
		fn   func(b *testing.B, ll *linkedlist.LinkedList[int])
	}{
		{name: "benchmarkLinkedListPopHead100k", fn: benchmarkLinkedListPop100k},
		{name: "benchmarkLinkedListPopHead1m", fn: benchmarkLinkedListPop1m},
		{name: "benchmarkLinkedListPopHead10m", fn: benchmarkLinkedListPop10m},
	}
	b.ResetTimer()
	for _, bb := range bc {
		b.Run(bb.name, func(b *testing.B) {
			bb.fn(b, ll)
		})
	}
}

func benchmarkLinkedListPopHead100k(b *testing.B, ll *linkedlist.LinkedList[int]) {
	for i := 0; i < 100000; i++ {
		ll.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ll.PopHead()
	}
}

func benchmarkLinkedListPopHead1m(b *testing.B, ll *linkedlist.LinkedList[int]) {
	for i := 0; i < 1000000; i++ {
		ll.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ll.PopHead()
	}
}

func benchmarkLinkedListPopHead10m(b *testing.B, ll *linkedlist.LinkedList[int]) {
	for i := 0; i < 10000000; i++ {
		ll.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ll.PopHead()
	}
}

func BenchmarkLinkedList_Find(b *testing.B) {
	ll := linkedlist.NewLinkedList[int]()

	bc := []struct {
		name string
		fn   func(b *testing.B, ll *linkedlist.LinkedList[int])
	}{
		{name: "benchmarkLinkedListFind100k", fn: benchmarkLinkedListPop100k},
		{name: "benchmarkLinkedListFind1m", fn: benchmarkLinkedListPop1m},
		{name: "benchmarkLinkedListFind10m", fn: benchmarkLinkedListPop10m},
	}
	b.ResetTimer()
	for _, bb := range bc {
		b.Run(bb.name, func(b *testing.B) {
			bb.fn(b, ll)
		})
	}
}

func benchmarkLinkedListFind100k(b *testing.B, ll *linkedlist.LinkedList[int]) {
	for i := 0; i < 100000; i++ {
		ll.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ll.Find(i)
	}
}

func benchmarkLinkedListFind1m(b *testing.B, ll *linkedlist.LinkedList[int]) {
	for i := 0; i < 1000000; i++ {
		ll.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ll.Find(i)
	}
}

func benchmarkLinkedListFind10m(b *testing.B, ll *linkedlist.LinkedList[int]) {
	for i := 0; i < 10000000; i++ {
		ll.Append(i)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ll.Find(i)
	}
}
