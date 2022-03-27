package queue_test

import (
	"testing"

	"github.com/HotPotatoC/sture/queue"
)

func BenchmarkQueue_Enqueue(b *testing.B) {
	bc := []struct {
		name string
		fn   func(*testing.B)
	}{
		{name: "Queue", fn: benchmarkQueueEnqueue},
		{name: "Priority Queue", fn: benchmarkPriorityQueueEnqueue},
		{name: "Ring Buffer", fn: benchmarkRingBufferEnqueue},
	}

	b.ResetTimer()
	for _, bb := range bc {
		b.Run(bb.name, func(b *testing.B) {
			bb.fn(b)
		})
	}
}

func benchmarkQueueEnqueue(b *testing.B) {
	q := queue.NewQueue[int]()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i)
	}
}

func benchmarkPriorityQueueEnqueue(b *testing.B) {
	q := queue.NewPriorityQueue[int]()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		q.Enqueue(i, i)
	}
}

func benchmarkRingBufferEnqueue(b *testing.B) {
	bc := []struct {
		name string
		cap  int
	}{
		{name: "1024", cap: 1024},         // 2^10
		{name: "65536", cap: 65536},       // 2^16
		{name: "262144", cap: 262144},     // 2^18
		{name: "1048576", cap: 1048576},   // 2^20
		{name: "16777216", cap: 16777216}, // 2^24
	}

	b.ResetTimer()
	for _, bb := range bc {
		b.Run(bb.name, func(b *testing.B) {
			q := queue.NewRingBuffer[int](bb.cap)

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				q.Enqueue(i)
			}
		})
	}
}

func BenchmarkQueue_Dequeue(b *testing.B) {
	bc := []struct {
		name string
		fn   func(*testing.B)
	}{
		{name: "Queue", fn: benchmarkQueueDequeue},
		{name: "Priority Queue", fn: benchmarkPriorityQueueDequeue},
		{name: "Ring Buffer", fn: benchmarkRingBufferDequeue},
	}

	b.ResetTimer()
	for _, bb := range bc {
		b.Run(bb.name, func(b *testing.B) {
			bb.fn(b)
		})
	}
}

func benchmarkQueueDequeue(b *testing.B) {
	bc := []struct {
		name string
		size int
	}{
		{name: "100k", size: 100000},
		{name: "1m", size: 1000000},
		{name: "10m", size: 10000000},
	}

	b.ResetTimer()
	for _, bb := range bc {
		b.Run(bb.name, func(b *testing.B) {
			q := queue.NewQueue[int]()

			for i := 0; i < bb.size; i++ {
				q.Enqueue(i)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				q.Dequeue()
			}
		})
	}
}

func benchmarkPriorityQueueDequeue(b *testing.B) {
	bc := []struct {
		name string
		size int
	}{
		{name: "100k", size: 100000},
		{name: "1m", size: 1000000},
		{name: "10m", size: 10000000},
	}

	b.ResetTimer()
	for _, bb := range bc {
		b.Run(bb.name, func(b *testing.B) {
			q := queue.NewPriorityQueue[int]()

			for i := 0; i < bb.size; i++ {
				q.Enqueue(i, i)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				q.Dequeue()
			}
		})
	}
}

func benchmarkRingBufferDequeue(b *testing.B) {
	bc := []struct {
		name string
		size int
	}{
		{name: "1024", size: 1024},         // 2^10
		{name: "65536", size: 65536},       // 2^16
		{name: "262144", size: 262144},     // 2^18
		{name: "1048576", size: 1048576},   // 2^20
		{name: "16777216", size: 16777216}, // 2^24
	}

	b.ResetTimer()
	for _, bb := range bc {
		b.Run(bb.name, func(b *testing.B) {
			q := queue.NewRingBuffer[int](bb.size)

			for i := 0; i < bb.size; i++ {
				q.Enqueue(i)
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				q.Dequeue()
			}
		})
	}
}
