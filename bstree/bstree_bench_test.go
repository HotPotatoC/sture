package bstree_test

import (
	"math"
	"math/rand"
	"testing"
	"time"

	"github.com/HotPotatoC/sture"
	"github.com/HotPotatoC/sture/bstree"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func randomIntRange(min, max int) int {
	return rand.Intn(max-min) + min
}

func BenchmarkBSTree_InsertRandom(b *testing.B) {
	bst := bstree.New(1, sture.Compare[int])

	for i := 0; i < b.N; i++ {
		_ = bst.Insert(bst.Root(), randomIntRange(1, 10000))
	}
}

func BenchmarkBSTree_SearchRandom_N(b *testing.B) {
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
			bst := bstree.New(randomIntRange(1, bb.size), sture.Compare[int])

			min := math.MaxInt

			for i := 0; i < bb.size; i++ {
				n := randomIntRange(1, bb.size)
				_ = bst.Insert(bst.Root(), n)

				if min > n {
					min = n
				}
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = bst.Search(bst.Root(), randomIntRange(min, bb.size))
			}
		})
	}
}

func BenchmarkBSTree_RemoveRandom_N(b *testing.B) {
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
			bst := bstree.New(randomIntRange(1, bb.size), sture.Compare[int])

			min := math.MaxInt

			for i := 0; i < bb.size; i++ {
				n := randomIntRange(1, bb.size)
				_ = bst.Insert(bst.Root(), n)

				if min > n {
					min = n
				}
			}

			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = bst.Remove(bst.Root(), randomIntRange(min, bb.size))
			}
		})
	}
}
