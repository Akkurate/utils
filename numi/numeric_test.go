package numi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelta(t *testing.T) {
	v := []int{0, 1, 3, 5, 5}
	resp := []int{1, 2, 2, 0}
	assert.Equal(t, resp, Delta(v))
	assert.Equal(t, 0, len(Delta([]int{1})))

}

func TestFind(t *testing.T) {
	v := []int{3, 4, 5, 6, 7, -1}

	f, _ := FindIndex(v, 5)
	assert.Equal(t, 2, f)
	f, _ = FindIndex(v, 999)
	assert.Equal(t, -1, f)

	c := Contains(v, -1)
	assert.Equal(t, true, c)
	c = Contains(v, 999)
	assert.Equal(t, false, c)
	v2 := []int{-1, 0, 1, 2, 3, 4, 5}
	expected := []int{1,2,3}
	assert.Equal(t, expected, FindMissingIndexes(v,v2))
}

func TestSliceOf(t *testing.T) {
	v := SliceOf(9, 9)
	assert.Equal(t, 9, len(v))
	v = Insert(v, 5, 1)
	assert.Equal(t, 10, len(v))
	assert.Equal(t, 1, v[5])
	v = RemoveFrom(v, 2)
	assert.Equal(t, 9, len(v))
	assert.Equal(t, 1, v[4])
	v = Remove(v, 9)
	assert.Equal(t, 8, len(v))
	assert.Equal(t, 1, v[3])
}

func TestCumsum(t *testing.T) {
	v := SliceOf(2, 10)
	r := Cumsum(v)
	expected := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	assert.Equal(t, expected, r)

}
func TestUnique(t *testing.T) {
	s := []int{1,2,3,3,3,4,5}
	expected := []int{1,2,3,4,5}
	assert.Equal(t, expected, Unique(s))
	assert.Equal(t, true, Contains(s, 3))
	assert.Equal(t, false, Contains(s,6))
	
	f, _ := FindIndex(s, 3)
	assert.Equal(t, 2, f)
	f, _ = FindIndex(s, 6)
	assert.Equal(t, -1, f)
}