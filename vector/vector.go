package vector

import "fmt"

// Vector is a linear times structure, the internal is a slice
type Vector[T any] struct {
	items []T
}

// New create a new Vector
func New[T any](size int) *Vector[T] {
	return &Vector[T]{items: make([]T, 0, size)}
}

// NewFrom create a new Vector and initialize it
func NewFrom[T any](items ...T) *Vector[T] {
	return &Vector[T]{items: items}
}

// Size returns the size of the vector
func (v *Vector[T]) Size() int {
	return len(v.items)
}

// Cap returns the capacity of the vector
func (v *Vector[T]) Cap() int {
	return cap(v.items)
}

// Empty returns true if the vector is empty, otherwise returns false
func (v *Vector[T]) Empty() bool {
	return len(v.items) == 0
}

// At returns the item at position pos, returns nil if pos is out off range
func (v *Vector[T]) At(pos int) T {
	if pos < 0 || pos >= v.Size() {
		panic("out off range")
	}

	return v.items[pos]
}

// Front return the first item in the vector, returns nil if the vector is empty
func (v *Vector[T]) Front() T {
	return v.At(0)
}

// Back return the last item in the vector, returns nil if the vector is empty
func (v *Vector[T]) Back() T {
	return v.At(v.Size() - 1)
}

// Push pushes item to the back of the vector
func (v *Vector[T]) Push(item T) {
	v.items = append(v.items, item)
}

// Pop returns the last item of the vector and erase it, returns nil if the vector is empty
func (v *Vector[T]) Pop() T {
	if v.Empty() {
		panic("out off range")
	}

	item := v.Back()
	v.items = v.items[:len(v.items)-1]

	return item
}

// Insert inserts item at position pos, returns nil if pos is out off range
func (v *Vector[T]) Insert(pos int, item T) {
	if pos < 0 || pos > v.Size() {
		panic("out off range")
	}

	v.Push(item)

	if pos != v.Size() {
		copy(v.items[pos+1:], v.items[pos:])
		v.items[pos] = item
	}
}

// Remove removes the item of the position pos in the vector, returns nil if pos is out off range
func (v *Vector[T]) Remove(pos int) T {
	if pos < 0 || pos >= v.Size() {
		panic("out off range")
	}

	item := v.At(pos)

	if pos != v.Size()-1 {
		copy(v.items[pos:], v.items[pos+1:])
	}

	v.items = v.items[:len(v.items)-1]

	return item
}

// Replace write the item of the position pos in the vector, returns nil if pos is out off
func (v *Vector[T]) Replace(pos int, item T) {
	if pos < 0 || pos >= v.Size() {
		panic("out off range")
	}

	v.items[pos] = item
}

// Clear clears all data in the vector
func (v *Vector[T]) Clear() {
	v.items = v.items[:0]
}

// Data returns internal data of the vector
func (v *Vector[T]) Data() []T {
	return v.items
}

// Clone returns a clone of the vector
func (v *Vector[T]) Clone() *Vector[T] {
	data := make([]T, v.Size(), v.Cap())
	copy(data, v.items)

	return &Vector[T]{items: data}
}

// String returns a string representation of the vector
func (v *Vector[T]) String() string {
	return fmt.Sprintf("%v", v.items)
}
