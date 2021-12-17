package numf

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumRange(t *testing.T) {

	r := NumRange(0, 10, 1)
	assert.Equal(t, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, r)
	r = NumRange(-10, -4, 2)
	assert.Equal(t, []float64{-10, -8, -6, -4}, r)
	r = NumRange(-10, -4, -2)
	assert.Equal(t, 0, len(r))
	r = NumRange(0, 4, 3)
	assert.Equal(t, []float64{0, 3}, r)
	r = NumRange(1, 1, 3)
	assert.Equal(t, []float64{1}, r)
	r = NumRange(8, 4, 3)
	assert.Equal(t, 0, len(r))
}

func TestSetDefault(t *testing.T) {
	s := SetDefault(3.3, 6.6)
	assert.Equal(t, 6.6, s)
	s = SetDefault(3.3, 0.0)
	assert.Equal(t, 3.3, s)
}

func TestSwitchIf(t *testing.T) {
	a, b := SwitchIf(true, 1.1, 2.2)
	assert.Equal(t, 2.2, a)
	assert.Equal(t, 1.1, b)
	a, b = SwitchIf(false, 1.1, 2.2)
	assert.Equal(t, 1.1, a)
	assert.Equal(t, 2.2, b)
}

func TestRoundTo(t *testing.T) {
	x := []float64{0, 12, -12, 0.0000449, 34545523, -1232.2}
	y5 := []float64{0, 12, -12, 0.0000449, 34546000, -1232.2}
	prefixes := []string{"", "", "", "u", "M", ""}
	y2 := []float64{0, 12, -12, 0.000045, 35000000, -1200}
	for i, v := range x {
		assert.Equal(t, y5[i], RoundTo(v, 5))
		assert.Equal(t, y2[i], RoundTo(v, 2))
		rounded := RoundWithPrefix(v, 5, "unit", "")
		assert.Equal(t, y5[i], rounded.Rawvalue)
		assert.Equal(t, prefixes[i], rounded.Prefix)
		rounded = RoundWithPrefix(v, 2, "%", "k")
		assert.Equal(t, y2[i], rounded.Rawvalue)
		assert.Equal(t, "", rounded.Prefix)
	}

	x = []float64{12.5, 0, 1000, 119.9, 120.1}
	n20 := []float64{20, 0, 1000, 120, 120}
	for i, v := range x {
		assert.Equal(t, n20[i], RoundToNearest(v, 20))

	}
}

func TestIsEqualSlice(t *testing.T) {
	x := []float64{0, math.NaN(), math.Inf(1)}
	y := []float64{0, math.NaN(), math.Inf(1)}
	assert.Equal(t, true, IsEqualSlice(x, y))
	y[2] = math.Inf(-1)
	assert.Equal(t, false, IsEqualSlice(x, y))
}

func TestIsSamesign(t *testing.T) {
	assert.Equal(t, true, IsSameSign(3, 4))
	assert.Equal(t, true, IsSameSign(-2, -8))
	assert.Equal(t, false, IsSameSign(-3, 4))
	assert.Equal(t, false, IsSameSign(3, -4))
	assert.Equal(t, true, IsSameSign(0, -4))
	assert.Equal(t, true, IsSameSign(3, 0))
	assert.Equal(t, true, IsSameSign(0, 0))
}
