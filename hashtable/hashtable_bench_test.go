package hashtable_test

import (
	"testing"

	"github.com/HotPotatoC/sture/hashtable"
)

func BenchmarkOpenAddressingHashTable_Set(b *testing.B) {
	ht := hashtable.NewHashTable[int, int](8)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ht.Set(i, i)
	}
}

func BenchmarkOpenAddressingHashTable_Get(b *testing.B) {
	ht := hashtable.NewHashTable[int, int](8)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ht.Set(i, i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = ht.Get(i)
	}
}

func BenchmarkOpenAddressingHashTable_Del(b *testing.B) {
	ht := hashtable.NewHashTable[int, int](8)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ht.Set(i, i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ht.Del(i)
	}
}

func BenchmarkChainedHashTable_Set(b *testing.B) {
	ht := hashtable.NewChainedHashTable[int, int](8)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ht.Set(i, i)
	}
}

func BenchmarkChainedHashTable_Get(b *testing.B) {
	ht := hashtable.NewChainedHashTable[int, int](8)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ht.Set(i, i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, _ = ht.Get(i)
	}
}

func BenchmarkChainedHashTable_Del(b *testing.B) {
	ht := hashtable.NewChainedHashTable[int, int](8)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ht.Set(i, i)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ht.Del(i)
	}
}

func BenchmarkStdlibMap_Set(b *testing.B) {
	m := make(map[int]int)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m[i] = i
	}
}

func BenchmarkStdlibMap_Get(b *testing.B) {
	m := make(map[int]int)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m[i] = i
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_ = m[i]
	}
}

func BenchmarkStdlibMap_Del(b *testing.B) {
	m := make(map[int]int)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		m[i] = i
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		delete(m, i)
	}
}