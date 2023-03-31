package queue

// Queue implements a FIFO queue data structure
type Queue[T any] struct {
	items []T
}

// New create a new FIFO queue where the items are stored in a plain slice
func New[T any](size int) *Queue[T] {
	return &Queue[T]{items: make([]T, 0, size)}
}

// NewFrom create a new FIFO queue and initialize it
func NewFrom[T any](items ...T) *Queue[T] {
	return &Queue[T]{items: items}
}

// Size returns the FIFO queue size
func (q *Queue[T]) Size() int {
	return len(q.items)
}

// Cap returns the FIFO queue cap
func (q *Queue[T]) Cap() int {
	return cap(q.items)
}

// Empty returns true if the FIFO queue is empty, otherwise returns false
func (q *Queue[T]) Empty() bool {
	return len(q.items) == 0
}

// Push inserts a new element at the top of the queue
func (q *Queue[T]) Push(item T) {
	q.items = append(q.items, item)
}

// Peek returns the top element of the queue without removing is
func (q *Queue[T]) Peek() T {
	if q.Empty() {
		panic("out off range")
	}

	return q.items[0]
}

// Pop retrieves and removes the top element pushed into the queue
func (q *Queue[T]) Pop() T {
	if q.Empty() {
		panic("out off range")
	}

	item := q.Peek()
	q.items = q.items[1:]

	return item
}
