package hashtable_test

import (
	"fmt"
	"testing"

	"github.com/HotPotatoC/sture/hashtable"
)

func TestChainedHashTable_SetGet(t *testing.T) {
	ht := hashtable.NewChainedHashTable[string, string](8)

	ht.Set("a", "A")
	ht.Set("b", "B")
	ht.Set("c", "C")
	ht.Set("d", "D")
	ht.Set("e", "E")

	if v, ok := ht.Get("a"); !ok || v != "A" {
		t.Errorf("\nExpected: %s\nGot: %s", "A", v)
	}

	if v, ok := ht.Get("b"); !ok || v != "B" {
		t.Errorf("\nExpected: %s\nGot: %s", "B", v)
	}

	if v, ok := ht.Get("c"); !ok || v != "C" {
		t.Errorf("\nExpected: %s\nGot: %s", "C", v)
	}

	if v, ok := ht.Get("d"); !ok || v != "D" {
		t.Errorf("\nExpected: %s\nGot: %s", "D", v)
	}

	if v, ok := ht.Get("e"); !ok || v != "E" {
		t.Errorf("\nExpected: %s\nGot: %s", "E", v)
	}

	if v, ok := ht.Get("f"); ok {
		t.Errorf("\nExpected: %s\nGot: %s", "", v)
	}
}

func TestChainedHashTable_Del(t *testing.T) {
	ht := hashtable.NewChainedHashTable[string, string](8)

	ht.Set("a", "A")
	ht.Set("b", "B")
	ht.Set("c", "C")
	ht.Set("d", "D")
	ht.Set("e", "E")

	ht.Del("a")
	ht.Del("b")
	ht.Del("c")
	ht.Del("d")
	ht.Del("e")

	if v, ok := ht.Get("a"); ok {
		t.Errorf("\nExpected: %s\nGot: %s", "", v)
	}

	if v, ok := ht.Get("b"); ok {
		t.Errorf("\nExpected: %s\nGot: %s", "", v)
	}

	if v, ok := ht.Get("c"); ok {
		t.Errorf("\nExpected: %s\nGot: %s", "", v)
	}

	if v, ok := ht.Get("d"); ok {
		t.Errorf("\nExpected: %s\nGot: %s", "", v)
	}

	if v, ok := ht.Get("e"); ok {
		t.Errorf("\nExpected: %s\nGot: %s", "", v)
	}
}

func TestChainedHashTable_Set10m(t *testing.T) {
	ht := hashtable.NewChainedHashTable[string, string](8)

	for i := 0; i < 10000000; i++ {
		ht.Set(fmt.Sprintf("k%d", i), fmt.Sprintf("v%d", i))
	}

	if ht.Size() != 10000000 {
		t.Errorf("\nExpected size: %d\nGot: %d", 10000000, ht.Size())
	}

	failures := 0
	for i := 0; i < 10000000; i++ {
		if v, ok := ht.Get(fmt.Sprintf("k%d", i)); !ok || v != fmt.Sprintf("v%d", i) {
			failures++
		}
	}

	if failures > 0 {
		t.Errorf("\nFailures: %d", failures)
	}
}
