package slices

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestShouldNopWhenNothingIsAdded(t *testing.T) {
	s := []int{1, 2, 3}
	s = expandSlice(s)
	assert.Equal(t, s, []int{1, 2, 3})
}

func TestShouldAddOneItemWhileInCapacity(t *testing.T) {
	s := make([]int, 1, 2)
	s[0] = 8
	s = expandSlice(s, 4)
	assert.Equal(t, s, []int{8, 4})
}

func TestShouldAddSeveralItemWhileInCapacity(t *testing.T) {
	s := make([]int, 1, 10)
	s[0] = 8
	s = expandSlice(s, 4, 5, 6, 7, 8)
	assert.Equal(t, s, []int{8, 4, 5, 6, 7, 8})
}

func TestShouldIncreaseCapacityWhenNecessary(t *testing.T) {
	s := make([]int, 1) // cap(s) == len(s) == 1
	s[0] = 8
	s = expandSlice(s, 4)

	assertHasGrown(s, 1, t)
	assert.Equal(t, s, []int{8, 4})
}

func TestShouldAddSeveralItemsAndIncreaseCapacity(t *testing.T) {
	s := []int{1, 2}
	s = expandSlice(s, 4, 5, 6)

	assertHasGrown(s, 2, t)
	assert.Equal(t, s, []int{1, 2, 4, 5, 6})
}

func assertHasGrown(s []int, originalLength int, t *testing.T) {
	if hasGrown := cap(s) >= originalLength && len(s) > originalLength; !hasGrown {
		t.Errorf("Slice has not grown as expected (len %d and cap %d but "+
			"should hold more than %d elements)", len(s), cap(s), originalLength)
	}
}
