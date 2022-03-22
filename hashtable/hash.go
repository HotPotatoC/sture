package hashtable

import (
	"unsafe"

	"github.com/cespare/xxhash/v2"
)

// HashFunction is a hash function type.
type HashFunction[K comparable] func(key K, size int) uint64

// DefaultHashFunction is the default hash function (using xxhash).
func DefaultHashFunction[K comparable](key K, size int) uint64 {
	// reference: https://github.com/tidwall/hashmap/blob/master/map.go#L46
	var sKey string
	switch any(key).(type) {
	case string:
		sKey = *(*string)(unsafe.Pointer(&key))
	default:
		sKey = *(*string)(unsafe.Pointer(&struct {
			data unsafe.Pointer
			len  int
		}{unsafe.Pointer(&key), int(unsafe.Sizeof(key))}))
	}
	return xxhash.Sum64String(sKey) % uint64(size)
}