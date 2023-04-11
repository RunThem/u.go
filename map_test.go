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
