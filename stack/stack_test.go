package stack

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase(t *testing.T) {
	arr := []int{10, 11, 12, 13, 14}
	st := New[int](10)

	assert.True(t, st.Empty())

	for _, k := range arr {
		st.Push(k)
	}

	assert.Equal(t, len(arr), st.Size())

	assert.Equal(t, arr[len(arr)-1], st.Peek())

	for i := len(arr) - 1; i >= 0; i-- {
		assert.Equal(t, arr[i], st.Pop())
	}
}

func TestReplace(t *testing.T) {
	arr := []int{10, 11, 12, 13, 14}
	st := New[int](10)

	for _, k := range arr {
		st.Push(k)
	}

	assert.Equal(t, 14, st.Replace(100))
	assert.Equal(t, 100, st.Peek())

	arr[len(arr)-1] = 100
	for i := len(arr) - 1; i >= 0; i-- {
		assert.Equal(t, arr[i], st.Pop())
	}

	st.Clear()
	assert.True(t, st.Empty())
}

func TestClone(t *testing.T) {
	arr := []int{10, 11, 12, 13, 14}
	st := New[int](10)

	for _, k := range arr {
		st.Push(k)
	}

	st_1 := st.Clone()

	assert.Equal(t, len(arr), st_1.Size())
	assert.Equal(t, arr[len(arr)-1], st_1.Peek())

	for i := len(arr) - 1; i >= 0; i-- {
		assert.Equal(t, arr[i], st_1.Pop())
	}
}
