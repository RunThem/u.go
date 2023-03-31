package vector

// Vector is a linear data structure, the internal is a slice
type Vector[T any] struct {
	data []T
}

// New create a new Vector
func New[T any](size int) *Vector[T] {
	return &Vector[T]{data: make([]T, 0, size)}
}

// New create a new Vector and initialize it
func NewFrom[T any](vals ...T) *Vector[T] {
	return &Vector[T]{data: vals}
}

// Size returns the size of the vector
func (v *Vector[T]) Size() int {
	return len(v.data)
}

// Cap returns the capacity of the vector
func (v *Vector[T]) Cap() int {
	return cap(v.data)
}

// Empty returns true if the vector is empty, otherwise returns false
func (v *Vector[T]) Empty() bool {
	return len(v.data) == 0
}

// At returns the value at position pos, returns nil if pos is out off range
func (v *Vector[T]) At(pos int) T {
	if pos < 0 || pos >= v.Size() {
		panic("out off range")
	}

	return v.data[pos]
}

// Front return the first value in the vector, returns nil if the vector is empty
func (v *Vector[T]) Front() T {
	return v.At(0)
}

// Back return the last value in the vector, returns nil if the vector is empty
func (v *Vector[T]) Back() T {
	return v.At(v.Size() - 1)
}

// Push pushes val to the back of the vector
func (v *Vector[T]) Push(val T) {
	v.data = append(v.data, val)
}

// Pop returns the last val of the vector and erase it, returns nil if the vector is empty
func (v *Vector[T]) Pop() T {
	if v.Empty() {
		panic("out off range")
	}

	val := v.Back()
	v.data = v.data[:len(v.data)-1]

	return val
}

// Insert insert val at position pos, returns nil if pos is out off range
func (v *Vector[T]) Insert(pos int, val T) {
	if pos < 0 || pos > v.Size() {
		panic("out off range")
	}

	v.Push(val)
	copy(v.data[pos+1:], v.data[pos:])
	v.data[pos] = val
}

// Remove remove the value of the position pos in the vector, returns nil if pos is out off range
func (v *Vector[T]) Remove(pos int) T {
	if pos < 0 || pos >= v.Size() {
		panic("out off range")
	}

	val := v.At(pos)

	if pos != v.Size()-1 {
		copy(v.data[pos:], v.data[pos+1:])
	}

	v.data = v.data[:len(v.data)-1]

	return val
}
