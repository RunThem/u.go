package u_go

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestKeys(t *testing.T) {
	m := map[int]string{1: "1", 2: "2", 3: "3"}

	a := Keys(m)

	assert.True(t, reflect.DeepEqual([]int{1, 2, 3}, a))
}

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5}

	Map(arr, func(i int) int {
		return i * 2
	})

	assert.True(t, reflect.DeepEqual([]int{2, 4, 6, 8, 10}, arr))
}
