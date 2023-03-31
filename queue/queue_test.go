package queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase(t *testing.T) {
	arr := []int{10, 11, 12, 13, 14}
	que := New[int](10)

	assert.True(t, que.Empty())

	for _, k := range arr {
		que.Push(k)
	}

	assert.Equal(t, len(arr), que.Size())
	assert.Equal(t, arr[0], que.Peek())

	for _, k := range arr {
		assert.Equal(t, k, que.Pop())
	}
}

func TestNewFrom(t *testing.T) {
	arr := []int{10, 11, 12, 13, 14}
	que := NewFrom(10, 11, 12, 13, 14)

	assert.Equal(t, len(arr), que.Size())
	for _, k := range arr {
		assert.Equal(t, k, que.Pop())
	}

	que = NewFrom(arr...)

	assert.Equal(t, len(arr), que.Size())
	for _, k := range arr {
		assert.Equal(t, k, que.Pop())
	}
}

func TestClone(t *testing.T) {
	arr := []int{10, 11, 12, 13, 14}
	que := NewFrom(10, 11, 12, 13, 14)

	que_1 := que.Clone()

	assert.Equal(t, len(arr), que_1.Size())
	for _, k := range arr {
		assert.Equal(t, k, que_1.Pop())
	}
}
