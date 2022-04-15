package tree_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/HotPotatoC/sture"
	"github.com/HotPotatoC/sture/tree"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func BenchmarkAVLTree_Insert(b *testing.B) {
	avl := tree.NewAVLTree[int, int](sture.Compare[int])

	for i := 0; i < b.N; i++ {
		avl.Insert(i, i)
	}
}

func BenchmarkAVLTree_SearchRandom_N(b *testing.B) {
	bc := []struct {
		name string
		size int
	}{
		{name: "100", size: 100},
		{name: "1k", size: 1000},
		{name: "10k", size: 10000},
		{name: "100k", size: 100000},
		{name: "1m", size: 1000000},
	}

	for _, bb := range bc {
		b.Run(bb.name, func(b *testing.B) {
			avl := tree.NewAVLTree[int, int](sture.Compare[int])

			min := math.MaxInt

			for i := 0; i < bb.size; i++ {
				n := randomIntRange(1, bb.size)
				avl.Insert(n, n)

				if min > n {
					min = n
				}
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = avl.Search(randomIntRange(min, bb.size))
			}
		})
	}
}

func BenchmarkAVLTree_RemoveRandom_N(b *testing.B) {
	bc := []struct {
		name string
		size int
	}{
		{name: "100", size: 100},
		{name: "1k", size: 1000},
		{name: "10k", size: 10000},
		{name: "100k", size: 100000},
		{name: "1m", size: 1000000},
	}

	for _, bb := range bc {
		b.Run(bb.name, func(b *testing.B) {
			avl := tree.NewAVLTree[int, int](sture.Compare[int])

			min := math.MaxInt

			for i := 0; i < bb.size; i++ {
				n := randomIntRange(1, bb.size)
				avl.Insert(n, n)

				if min > n {
					min = n
				}
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				avl.Remove(randomIntRange(min, bb.size))
			}
		})
	}
}
