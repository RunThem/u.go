package u_go

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilter(t *testing.T) {
	M := assert.New(t)

	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	M.Equal([]int{2, 4, 6, 8, 10}, Filter(items, func(item int) bool {
		return item%2 == 0
	}))

	M.Equal([]int{1, 3, 5, 7, 9}, Filter(items, func(item int) bool {
		return item%2 != 0
	}))

	items1 := map[int]string{1: "john", 2: "doe", 3: "fred"}
	M.Equal(map[int]string{1: "john"}, FilterMap(items1, func(s string) bool {
		return s == "john"
	}))
}
