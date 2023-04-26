package list

// Element is an element of a linked list
type Element[T any] struct {
	next, prev *Element[T]
	Value      T
}

// Next returns the next list element or nil
func (e *Element[T]) Next() *Element[T] {
	if e.next != nil {
		return e.next
	}

	return nil
}

// Prev returns the previous list element or nil
func (e *Element[T]) Prev() *Element[T] {
	if e.prev != nil {
		return e.prev
	}

	return nil
}

// List represents a doubly linked list.
// The zero value for List is an empty list ready to use
type List[T any] struct {
	head *Element[T]
	tail *Element[T]
	len  int // current list length excluding (this) sentinel element
}

// New returns an initialized list
func New[T any]() *List[T] {
	return &List[T]{}
}

func NewFrom[T any](items ...T) *List[T] {
	l := New[T]()

	for _, it := range items {
		l.PushBack(it)
	}

	return l
}

// Size returns the number of elements of list l
// The complexity is O(1).
func (l *List[T]) Size() int {
	return l.len
}

// Front returns the first element of list l or nil if the list is empty
func (l *List[T]) Front() *Element[T] {
	if l.len == 0 {
		return nil
	}

	return l.head
}

// Back returns the last element of list l or nil if the list is empty
func (l *List[T]) Back() *Element[T] {
	if l.len == 0 {
		return nil
	}

	return l.tail
}

// Remove removes e from l if e is an element of list l.
// It returns the element value e.Value.
// The element must not be nil
func (l *List[T]) Remove(e *Element[T]) T {
	if e == l.head {
		l.head = e.next
		e.next = nil
	} else if e == l.tail {
		l.tail = e.prev
		e.prev = nil
	} else {
		prev := e.prev
		next := e.next

		prev.next = next
		next.prev = prev

		e.prev = nil
		e.next = nil
	}

	return e.Value
}

// PushFront inserts a new element e with value v at the front of list l
func (l *List[T]) PushFront(v T) {
	e := &Element[T]{Value: v}

	if l.len == 0 {
		l.head = e
		l.tail = e
	} else {
		l.head.prev = e
		e.next = l.head
		l.head = e
	}

	l.len++
}

// PushBack inserts a new element e with value v at the back of list l
func (l *List[T]) PushBack(v T) {
	e := &Element[T]{Value: v}

	if l.len == 0 {
		l.head = e
		l.tail = e
	} else {
		l.tail.next = e
		e.prev = l.tail
		l.tail = e
	}

	l.len++
}

// PopFront returns the first element in the list l
func (l *List[T]) PopFront() *Element[T] {
	var e *Element[T]

	if l.len == 0 {
		return nil
	}

	if l.len == 1 {
		e = l.head
		l.head = nil
		l.tail = nil
	} else {
		e = l.head
		l.head = e.next
	}

	l.len--
	e.prev = nil
	e.next = nil

	return e
}

// PopBack returns the last element in the list l
func (l *List[T]) PopBack() *Element[T] {
	var e *Element[T]

	if l.len == 0 {
		return nil
	}

	if l.len == 1 {
		e = l.tail
		l.head = nil
		l.tail = nil
	} else {
		e = l.tail
		l.tail = e.prev
	}

	l.len--
	e.prev = nil
	e.next = nil

	return e
}

func (l *List[T]) InsertBefore(e *Element[T], v T) {
	elem := &Element[T]{Value: v}

	if e == l.head {
		e.prev = elem
		l.head = elem
	} else {
		prev := e.prev

		elem.prev = prev
		elem.next = e

		prev.next = elem
		e.prev = elem
	}

	l.len++
}

func (l *List[T]) InsertAfter(e *Element[T], v T) {
	elem := &Element[T]{Value: v}

	if e == l.tail {
		e.next = elem
		l.tail = elem
	} else {
		next := e.next

		elem.prev = e
		elem.next = next

		e.next = elem
		next.prev = elem
	}

	l.len++
}

// Find returns the index of the first found element
func (l *List[T]) Find(fn func(T) bool) *Element[T] {
	for it := l.head; it != nil; it = it.next {
		if fn(it.Value) {
			return it
		}
	}

	return nil
}

// Filter returns all the elements from the collection which satisfies the conditional logic of callback function
func (l *List[T]) Filter(fn func(T) bool) *List[T] {
	lst := New[T]()

	for it := l.head; it != nil; it = it.next {
		if fn(it.Value) {
			lst.PushBack(it.Value)
		}
	}

	return lst
}

// Map maps an element to another value
func (l *List[T]) Map(fn func(T) T) {
	for it := l.head; it != nil; it = it.next {
		it.Value = fn(it.Value)
	}
}

// MapBy maps an element to another value
func (l *List[T]) MapBy(fn func(*T)) {
	for it := l.head; it != nil; it = it.next {
		fn(&it.Value)
	}
}
