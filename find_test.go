package u_go

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestFind(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(t, 0, Find(arr, 1))
	assert.Equal(t, 4, Find(arr, 5))
	assert.Equal(t, 9, Find(arr, 10))
	assert.Equal(t, -1, Find(arr, 11))
}

func TestFindBy(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(t, 0, FindBy(arr, func(it int) bool {
		return it == 1
	}))
	assert.Equal(t, 4, FindBy(arr, func(it int) bool {
		return it == 5
	}))
	assert.Equal(t, 9, FindBy(arr, func(it int) bool {
		return it == 10
	}))
	assert.Equal(t, -1, FindBy(arr, func(it int) bool {
		return it == 11
	}))
}

func TestFindAll(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	m := FindAll(arr, func(it int) bool {
		return it%2 == 0
	})

	mm := map[int]int{1: 2, 3: 4, 5: 6, 7: 8, 9: 10}
	assert.True(t, reflect.DeepEqual(m, mm))
}

func TestMin(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(t, 0, Min(arr))
}

func TestMax(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(t, 9, Max(arr))
}

func TestMinBy(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(t, 0, MinBy(arr, func(a int, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}

		return 0
	}))
}

func TestMaxBy(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	assert.Equal(t, 9, MaxBy(arr, func(a int, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}

		return 0
	}))
}
