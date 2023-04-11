package u_go

import "golang.org/x/exp/constraints"

// Find returns the index of the first found element
func Find[T constraints.Ordered](s []T, elem T) int {
	for k, v := range s {
		if v == elem {
			return k
		}
	}

	return -1
}

// FindBy returns the index of the first found element
func FindBy[T any](s []T, fn func(T) bool) int {
	for k, v := range s {
		if fn(v) {
			return k
		}
	}

	return -1
}

// FindAll returns a map whose key is the index of the element that meets the criteria and value is
// the element that meets the criteria
func FindAll[T any](s []T, fn func(T) bool) map[int]T {
	m := make(map[int]T, len(s))

	for k, v := range s {
		if fn(v) {
			m[k] = v
		}
	}

	return m
}

// Min returns the index of the minimum value in slice
func Min[T constraints.Ordered](s []T) int {
	if len(s) == 0 {
		return -1
	}

	min := 0
	for i := 1; i < len(s); i++ {
		if s[i] < s[min] {
			min = i
		}
	}

	return min
}

// MinBy finds the minimum value of a slice
// fn first element greater than the second element returns 1, equality returns 0, and less than -1
func MinBy[T any](s []T, fn func(T, T) int) int {
	if len(s) == 0 {
		return -1
	}

	min := 0
	for i := 1; i < len(s); i++ {
		if fn(s[i], s[min]) == -1 {
			min = i
		}
	}

	return min
}

// Max returns the index of the maximum value in slice
func Max[T constraints.Ordered](s []T) int {
	if len(s) == 0 {
		return -1
	}

	max := 0
	for i := 1; i < len(s); i++ {
		if s[i] > s[max] {
			max = i
		}
	}

	return max
}

// MaxBy returns the index of the maximum value in slice
// fn first element greater than the second element returns 1, equality returns 0, and less than -1
func MaxBy[T any](s []T, fn func(T, T) int) int {
	if len(s) == 0 {
		return -1
	}

	max := 0
	for i := 1; i < len(s); i++ {
		if fn(s[i], s[max]) == 1 {
			max = i
		}
	}

	return max
}
