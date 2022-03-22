package hashtable

const (
	loadFactor = 0.85
)

// HashTable is a hash table.
type HashTable[K comparable, V any] struct {
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
	// buckets is the array of entries.
	buckets []*Entry[K, V]
	// hashFunc is the hashing function.
	hashFunc HashFunction[K]
}

// NewHashTable returns a new hash table with linear probing.
func NewHashTable[K comparable, V any](cap int, hashFunc ...HashFunction[K]) *HashTable[K, V] {
	ht := new(HashTable[K, V])

	if hashFunc == nil {
		// use default hash function if not specified by the caller
		ht.hashFunc = DefaultHashFunction[K]
	} else {
		ht.hashFunc = hashFunc[0]
	}

	c := 8
	for c < cap {
		c *= 2
	}

	ht.cap = c
	ht.buckets = make([]*Entry[K, V], c)
	ht.growAt = int(float64(len(ht.buckets)) * loadFactor)
	ht.shrinkAt = int(float64(len(ht.buckets)) * (1 - loadFactor))

	return ht
}

// resize resizes the hash table.
func (ht *HashTable[K, V]) resize(newCap int) {
	newht := NewHashTable[K, V](newCap)

	for _, entry := range ht.buckets {
		if entry != nil {
			newht.set(entry.key, entry.value)
		}
	}

	*ht = *newht
}

// Set sets the value of the key.
func (ht *HashTable[K, V]) Set(key K, value V) {
	if ht.size >= ht.growAt {
		ht.resize(len(ht.buckets) * 2)
	}

	ht.set(key, value)
}

func (ht *HashTable[K, V]) set(key K, value V) {
	probe := ht.hashFunc(key, len(ht.buckets))
	entry := NewEntry(key, value)

	for {
		if ht.buckets[probe] == nil {
			ht.buckets[probe] = entry
			ht.size++
			return
		}

		if ht.buckets[probe].key == key {
			ht.buckets[probe].value = value
			return
		}

		probe = (probe + 1) % uint64(ht.cap)
	}

}

// Get gets the value of the key.
func (ht *HashTable[K, V]) Get(key K) (V, bool) {
	return ht.get(key)
}

func (ht *HashTable[K, V]) get(key K) (V, bool) {
	probe := ht.hashFunc(key, len(ht.buckets))

	if ht.buckets[probe] == nil {
		return *new(V), false
	}

	for i := 0; i < ht.cap; i++ {
		if ht.buckets[probe] != nil && ht.buckets[probe].key == key {
			return ht.buckets[probe].value, true
		}

		probe = (probe + 1) % uint64(ht.cap)
	}

	return *new(V), false
}

// Del deletes the key.
func (ht *HashTable[K, V]) Del(key K) bool {
	if ht.size <= ht.shrinkAt {
		ht.resize(len(ht.buckets) / 2)
	}

	return ht.del(key)
}

func (ht *HashTable[K, V]) del(key K) bool {
	probe := ht.hashFunc(key, len(ht.buckets))

	if ht.buckets[probe] == nil {
		return false
	}

	for i := 0; i < ht.cap; i++ {
		if ht.buckets[probe] != nil && ht.buckets[probe].key == key {
			ht.buckets[probe] = nil
			return true
		}

		probe = (probe + 1) % uint64(ht.cap)
	}

	return false
}

// Exists returns true if the key exists.
func (ht *HashTable[K, V]) Exists(key K) bool {
	_, ok := ht.get(key)
	return ok
}

// Size returns the number of elements in the hash table.
func (ht *HashTable[K, V]) Size() int {
	return ht.size
}

// Cap returns the capacity of the hash table.
func (ht *HashTable[K, V]) Cap() int {
	return ht.cap
}
