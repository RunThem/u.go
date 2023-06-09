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

	vec.Push(12, 15, 22)
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

func TestInsert(t *testing.T) {
	arr := []int{11, 12, 13, 14, 15}
	vec := NewFrom(12, 14)

	vec.Insert(1, 13)
	vec.Insert(0, 11)
	vec.Insert(4, 15)

	assert.Equal(t, len(arr), vec.Size())

	for k, v := range arr {
		assert.Equal(t, v, vec.At(k))
	}
}

func TestRemove(t *testing.T) {
	arr := []int{12, 14}
	vec := NewFrom(11, 12, 13, 14, 15)

	assert.Equal(t, 11, vec.Remove(0))
	assert.Equal(t, 13, vec.Remove(1))
	assert.Equal(t, 15, vec.Remove(2))

	assert.Equal(t, len(arr), vec.Size())

	for k, v := range arr {
		assert.Equal(t, v, vec.At(k))
	}
}

func TestReplace(t *testing.T) {
	arr := []int{15, 14, 13, 12, 11}
	vec := NewFrom(11, 12, 13, 14, 15)

	for k, v := range arr {
		vec.Replace(k, v)
	}

	for k, v := range arr {
		assert.Equal(t, v, vec.At(k))
	}
}

func TestClear(t *testing.T) {
	vec := NewFrom(11, 12, 13, 14, 15)

	vec.Clear()

	assert.Equal(t, 0, vec.Size())
}

func TestData(t *testing.T) {
	vec := NewFrom(11, 12, 13, 14, 15)

	v := vec.Data()

	assert.Equal(t, vec.Size(), len(v))

	for i := 0; i < len(v); i++ {
		v[i] = 100
	}

	for i := 0; i < vec.Size(); i++ {
		assert.Equal(t, 100, vec.At(i))
	}
}

func TestClone(t *testing.T) {
	vec := NewFrom(11, 12, 13, 14, 15)
	vec_1 := vec.Clone()

	assert.Equal(t, vec.Size(), vec_1.Size())

	for i := 0; i < vec.Size(); i++ {
		assert.Equal(t, vec.At(i), vec_1.At(i))
	}
}

func TestFind(t *testing.T) {
	vec := NewFrom(11, 12, 13, 14, 15)

	i := vec.Find(func(i int) bool {
		return i == 12
	})

	assert.Equal(t, 1, i)
}

func TestFilter(t *testing.T) {
	vec := NewFrom(11, 12, 13, 14, 15)

	vec_1 := vec.Filter(func(i int) bool {
		return i%2 == 0
	})

	assert.Equal(t, 2, vec_1.Size())
	assert.Equal(t, 12, vec_1.At(0))
	assert.Equal(t, 14, vec_1.At(1))
}

func TestMap(t *testing.T) {
	arr := []int{15, 14, 13, 12, 11}
	vec := New[int](len(arr))

	for _, v := range arr {
		vec.Push(v)
	}

	vec.Map(func(i int) int {
		return i * 2
	})

	assert.Equal(t, len(arr), vec.Size())

	for i, v := range arr {
		assert.Equal(t, v*2, vec.At(i))
	}
}
