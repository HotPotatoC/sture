package hashtable

// Entry is an entry in a hash table.
type Entry[K comparable, V any] struct {
	key   K
	value V

	next *Entry[K, V]
	prev *Entry[K, V]
}

// NewEntry returns a new entry.
func NewEntry[K comparable, V any](key K, value V) *Entry[K, V] {
	return &Entry[K, V]{
		key:   key,
		value: value,
	}
}

// Key returns the key of the entry.
func (e *Entry[K, V]) Key() K {
	return e.key
}

// Value returns the value of the entry.
func (e *Entry[K, V]) Value() V {
	return e.value
}

// Next returns the next entry in the list.
func (e *Entry[K, V]) Next() *Entry[K, V] {
	return e.next
}

// Prev returns the previous entry in the list.
func (e *Entry[K, V]) Prev() *Entry[K, V] {
	return e.prev
}