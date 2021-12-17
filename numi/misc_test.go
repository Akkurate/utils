package numi

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumRange(t *testing.T) {

	r := NumRange(0, 10, 1)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, r)
	r = NumRange(-10, -4, 2)
	assert.Equal(t, []int{-10, -8, -6, -4}, r)
	r = NumRange(-10, -4, -2)
	assert.Equal(t, 0, len(r))
	r = NumRange(0, 4, 3)
	assert.Equal(t, []int{0, 3}, r)
	r = NumRange(1, 1, 3)
	assert.Equal(t, []int{1}, r)
	r = NumRange(8, 4, 3)
	assert.Equal(t, 0, len(r))
}
func TestSetDefault(t *testing.T) {
	s := SetDefault(3, 6)
	assert.Equal(t, 6, s)
	s = SetDefault(3, 0.0)
	assert.Equal(t, 3, s)
}
func TestSwitchIf(t *testing.T) {
	a, b := SwitchIf(true, 1, 2)
	assert.Equal(t, 2, a)
	assert.Equal(t, 1, b)
	a, b = SwitchIf(false, 1, 2)
	assert.Equal(t, 1, a)
	assert.Equal(t, 2, b)
}