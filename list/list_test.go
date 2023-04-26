package list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBase(t *testing.T) {
	arr := []int{10, 11, 12, 13, 14, 15, 16, 17, 18}
	lst := New[int]()

	lst.PushBack(13)
	lst.PushBack(14)
	lst.PushBack(15)
	lst.PushBack(16)
	lst.PushBack(17)
	lst.PushBack(18)

	lst.PushFront(12)
	lst.PushFront(11)
	lst.PushFront(10)

	assert.Equal(t, len(arr), lst.Size())

	for i, it := 0, lst.Front(); it != nil; it = it.Next() {
		assert.Equal(t, arr[i], it.Value)
		i++
	}
}

func TestInsert(t *testing.T) {
	lst := New[int]()

	lst.PushBack(15)

	n := lst.Back()

	lst.InsertBefore(n, 11)
	lst.InsertBefore(n, 12)
	lst.InsertBefore(n, 13)
	lst.InsertBefore(n, 14)

	lst.InsertAfter(n, 18)
	lst.InsertAfter(n, 17)
	lst.InsertAfter(n, 16)

	arr := []int{11, 12, 13, 14, 15, 16, 17, 18}
	for i, it := 0, lst.Front(); it != nil; it = it.Next() {
		assert.Equal(t, arr[i], it.Value)
		i++
	}
}

func TestFind(t *testing.T) {
	arr := []int{10, 11, 12, 13, 14, 15, 16, 17, 18}

	lst := NewFrom(arr...)

	for i, it := 0, lst.Front(); it != nil; it = it.Next() {
		assert.Equal(t, arr[i], it.Value)
		i++
	}

	n := lst.Find(func(i int) bool {
		return i == 14
	})

	assert.Equal(t, 14, n.Value)
	assert.Equal(t, 13, n.prev.Value)
	assert.Equal(t, 15, n.next.Value)
}

func TestFilter(t *testing.T) {
	arr := []int{10, 11, 12, 13, 14, 15, 16, 17, 18}
	arr_1 := []int{10, 12, 14, 16, 18}

	lst := NewFrom(arr...)

	lst_1 := lst.Filter(func(i int) bool {
		return i%2 == 0
	})

	assert.Equal(t, len(arr_1), lst_1.Size())

	for i, it := 0, lst_1.Front(); it != nil; it = it.Next() {
		assert.Equal(t, arr_1[i], it.Value)
		i++
	}
}

func TestMap(t *testing.T) {
	arr := []int{10, 11, 12, 13, 14, 15, 16, 17, 18}

	lst := NewFrom(arr...)

	lst.Map(func(i int) int {
		return i * 2
	})

	for i, it := 0, lst.Front(); it != nil; it = it.Next() {
		assert.Equal(t, arr[i]*2, it.Value)
		i++
	}
}
