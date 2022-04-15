package sture

import "golang.org/x/exp/constraints"

// Compare[T constraints.Ordered] is a function that returns -1, 0, or 1 depending on whether x < y, x == y, or x > y.
func Compare[T constraints.Ordered](x, y T) int {
	if x < y {
		return -1
	} else if x > y {
		return 1
	}

	return 0
}

// Max returns the maximum of two values.
func Max[T constraints.Ordered](x, y T) T {
	if x > y {
		return x
	}

	return y
}

// Min returns the minimum of two values.
func Min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}

	return y
}
