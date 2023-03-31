package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
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
