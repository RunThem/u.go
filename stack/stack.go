package stack

// Stack implements the LIFO Stack
type Stack[T any] struct {
	items []T
}

// New creates a new LIFO stack where the data are stored in a plain slice
func New[T any](size int) *Stack[T] {
	return &Stack[T]{items: make([]T, 0, size)}
}

// Size returns the LIFO stack size
func (s *Stack[T]) Size() int {
	return len(s.items)
}

// Cap returns the LIFO stack cap
func (s *Stack[T]) Cap() int {
	return cap(s.items)
}

// Empty returns true if the LIFO stack is empty, otherwise returns false
func (s *Stack[T]) Empty() bool {
	return len(s.items) == 0
}

// Push inserts a new element at the top of the stack
func (s *Stack[T]) Push(item T) {
	s.items = append(s.items, item)
}

// Peek returns the top element of the stack without removing is
func (s *Stack[T]) Peek() T {
	if s.Empty() {
		panic("out off range")
	}

	return s.items[s.Size()-1]
}

// Pop retrieves and removes the top element pushed into the stack
func (s *Stack[T]) Pop() T {
	if s.Empty() {
		panic("out off range")
	}

	item := s.Peek()
	s.items = s.items[:len(s.items)-1]

	return item
}

// Replace replaces the emelent at the top of the stack
func (s *Stack[T]) Replace(item T) T {
	v := s.Peek()

	s.items[s.Size()-1] = item

	return v
}
