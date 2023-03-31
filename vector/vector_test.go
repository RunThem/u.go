package vector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBase(t *testing.T) {
	vec := New[int](10)

	assert.True(t, vec.Empty())
	assert.Equal(t, 0, vec.Size())
	assert.Equal(t, 10, vec.Cap())

	vec.Push(12)
	vec.Push(15)
	vec.Push(22)
	assert.False(t, vec.Empty())
	assert.Equal(t, 3, vec.Size())

	assert.Equal(t, 12, vec.Front())
	assert.Equal(t, 22, vec.Back())

	assert.Equal(t, 12, vec.At(0))
	assert.Equal(t, 15, vec.At(1))
	assert.Equal(t, 22, vec.At(2))

	assert.Equal(t, 22, vec.Pop())
	assert.Equal(t, 15, vec.Pop())
	assert.Equal(t, 1, vec.Size())
}

func TestNewFrom(t *testing.T) {
	arr := []int{12, 13, 14, 15}

	vec := NewFrom(12, 13, 14, 15)
	assert.Equal(t, 4, vec.Size())

	for k, v := range arr {
		assert.Equal(t, v, vec.At(k))
	}

	vec = NewFrom(arr...)
	assert.Equal(t, 4, vec.Size())

	for k, v := range arr {
		assert.Equal(t, v, vec.At(k))
	}
}
