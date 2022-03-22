package hashtable

const (
	// load factor for resizing (cht stands for chained hash table to
	// avoid redefining in hashtable.go)
	chtLoadFactor = 0.85
)

// bucket is a bucket in a chained hash table.
type bucket[K comparable, V any] struct {
	head *Entry[K, V] // head of the doubly linked list
	tail *Entry[K, V] // tail of the doubly linked list
}

// ChainedHashTable is a chained hash table.
type ChainedHashTable[K comparable, V any] struct {
	// cap is the capacity of the table.
	cap int
	// size is the number of entries in the table.
	size int
	// growAt this size (n buckets * load factor) the table will grow with
	// a new capacity of cap * 2
	growAt int
	// shrinkAt this size (n buckets * (1 - load factor)) the table will shrink
	// with a new capacity of cap / 2
	shrinkAt int
	// buckets is the array of buckets.
	buckets []*bucket[K, V]
	// hashFunc is the hashing function.
	hashFunc HashFunction[K]
}

// NewChainedHashTable returns a new chained hash table.
func NewChainedHashTable[K comparable, V any](cap int, hashFunc ...HashFunction[K]) *ChainedHashTable[K, V] {
	var hf HashFunction[K]

	if hashFunc == nil {
		// use default hash function if not specified by the caller
		hf = DefaultHashFunction[K]
	} else {
		hf = hashFunc[0]
	}

	c := 8
	for c < cap {
		c *= 2
	}

	cht := &ChainedHashTable[K, V]{
		buckets:  make([]*bucket[K, V], c),
		hashFunc: hf,
		cap:      c,
	}

	cht.growAt = int(float64(len(cht.buckets)) * chtLoadFactor)
	cht.shrinkAt = int(float64(len(cht.buckets)) * (1 - chtLoadFactor))

	return cht
}

// resize resizes the chained hash table.
func (cht *ChainedHashTable[K, V]) resize(newCap int) {
	ht := NewChainedHashTable[K, V](newCap)

	for _, bucket := range cht.buckets {
		for bucket != nil && bucket.head != nil {
			entry := bucket.head
			bucket.head = entry.next
			ht.set(entry.key, entry.value)
		}
	}

	*cht = *ht
}

// Set sets the value of the key.
func (cht *ChainedHashTable[K, V]) Set(key K, value V) {
	if cht.size >= cht.growAt {
		cht.resize(len(cht.buckets) * 2)
	}

	cht.set(key, value)
}

// set sets the value of the key in the chained hash table.
func (cht *ChainedHashTable[K, V]) set(key K, value V) {
	idx := cht.hashFunc(key, len(cht.buckets))
	newEntry := NewEntry(key, value)

	if cht.buckets[idx] == nil {
		cht.buckets[idx] = &bucket[K, V]{
			head: newEntry,
			tail: newEntry,
		}
		cht.size++
		return
	}

	cht.buckets[idx].tail.next = newEntry
	newEntry.prev = cht.buckets[idx].tail
	cht.buckets[idx].tail = newEntry
	cht.size++
}

// Get gets the value of the key.
func (cht *ChainedHashTable[K, V]) Get(key K) (V, bool) {
	return cht.get(key)
}

// get returns the value of the key in the chained hash table.
func (cht *ChainedHashTable[K, V]) get(key K) (V, bool) {
	if cht.size == 0 {
		return *new(V), false
	}

	idx := cht.hashFunc(key, len(cht.buckets))

	if cht.buckets[idx] == nil {
		return *new(V), false // *new(V) zero value of V
	}

	if cht.buckets[idx].head == nil || cht.buckets[idx].tail == nil {
		return *new(V), false // *new(V) zero value of V
	}

	if cht.buckets[idx].head.key == key {
		return cht.buckets[idx].head.value, true
	}

	if cht.buckets[idx].tail.key == key {
		return cht.buckets[idx].tail.value, true
	}

	for curr := cht.buckets[idx].head; curr != nil; curr = curr.next {
		if curr.key == key {
			return curr.value, true
		}
	}

	return *new(V), false
}

// Del deletes the key.
func (cht *ChainedHashTable[K, V]) Del(key K) bool {
	if cht.size <= cht.shrinkAt {
		cht.resize(len(cht.buckets) / 2)
	}

	return cht.del(key)
}

// del deletes the key from the chained hash table.
func (cht *ChainedHashTable[K, V]) del(key K) bool {
	if cht.size == 0 {
		return false
	}

	idx := cht.hashFunc(key, len(cht.buckets))

	if cht.buckets[idx] == nil {
		return false
	}

	if cht.buckets[idx].head == nil || cht.buckets[idx].tail == nil {
		return false // *new(V) zero value of V
	}

	if cht.buckets[idx].head.key == key {
		if cht.buckets[idx].head.next == nil {
			cht.buckets[idx] = nil
			cht.size--
			return true
		}
		cht.buckets[idx].head = cht.buckets[idx].head.next
		cht.buckets[idx].head.prev = nil
		return true
	}

	if cht.buckets[idx].tail.key == key {
		cht.buckets[idx].tail = cht.buckets[idx].tail.prev
		cht.buckets[idx].tail.next = nil
		return true
	}

	for curr := cht.buckets[idx].head; curr != nil; curr = curr.next {
		if curr.key == key {
			curr.prev.next = curr.next
			curr.next.prev = curr.prev
			return true
		}
	}

	return false
}

// Exists returns true if the key exists.
func (cht *ChainedHashTable[K, V]) Exists(key K) bool {
	_, ok := cht.get(key)
	return ok
}

// Size returns the number of elements in the hash table.
func (cht *ChainedHashTable[K, V]) Size() int {
	return cht.size
}

// Cap returns the capacity of the hash table.
func (cht *ChainedHashTable[K, V]) Cap() int {
	return len(cht.buckets)
}
