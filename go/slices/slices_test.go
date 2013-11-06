package slices

import (
	"testing"
)

func TestShouldNopWhenNothingIsAdded(t *testing.T) {
	s := []int{1, 2, 3}
	s = expandSlice(s)
	assertEqual(s, []int{1, 2, 3}, t)
}

func TestShouldAddOneItemWhileInCapacity(t *testing.T) {
	s := make([]int, 1, 2)
	s[0] = 8
	s = expandSlice(s, 4)
	assertEqual(s, []int{8, 4}, t)
}

func TestShouldAddSeveralItemWhileInCapacity(t *testing.T) {
	s := make([]int, 1, 10)
	s[0] = 8
	s = expandSlice(s, 4, 5, 6, 7, 8)
	assertEqual(s, []int{8, 4, 5, 6, 7, 8}, t)
}

func TestShouldIncreaseCapacityWhenNecessary(t *testing.T) {
	s := make([]int, 1) // cap(s) == len(s) == 1
	s[0] = 8
	s = expandSlice(s, 4)

	assertHasGrown(s, 1, t)
	assertEqual(s, []int{8, 4}, t)
}

func TestShouldAddSeveralItemsAndIncreaseCapacity(t *testing.T) {
	s := []int{1, 2}
	s = expandSlice(s, 4, 5, 6)

	assertHasGrown(s, 2, t)
	assertEqual(s, []int{1, 2, 4, 5, 6}, t)
}

func assertHasGrown(s []int, originalLength int, t *testing.T) {
	if hasGrown := cap(s) >= len(s) && len(s) > 1; !hasGrown {
		t.Errorf("Slice has not grown as expected (len %d and cap %d but "+
			"should hold 2 elements)", len(s), cap(s))
	}
}

func assertEqual(s1, s2 []int, t *testing.T) {
	if len(s1) != len(s2) {
		t.Errorf("Not same length: %v and %v", s1, s2)
	}

	for v := range s1 {
		if s1[v] != s2[v] {
			t.Errorf("Differing item at position %d : %v versus %v\n%v\t%v",
				v, s1[v], s2[v], s1, s2)
		}
	}
}
