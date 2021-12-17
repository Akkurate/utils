package numf

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelta(t *testing.T) {
	v := []float64{0, 1, 3, 5, 5}
	resp := []float64{1, 2, 2, 0}
	assert.Equal(t, resp, Delta(v))
	assert.Equal(t, 0, len(Delta([]float64{1})))

}

func TestCompare(t *testing.T) {

	a := []float64{3, 9, math.NaN(), math.Inf(1), 4, math.Inf(1)}
	b := []float64{9, 1, 3, math.Inf(-1), math.Inf(1), math.NaN()}
	maxes := []float64{9, 9, 3, math.Inf(1), math.Inf(1), math.Inf(1)}
	mins := []float64{3, 1, 3, math.Inf(-1), 4, math.Inf(1)}
	for i, v := range a {
		max, min := Compare(v, b[i])
		assert.Equal(t, maxes[i], max)
		assert.Equal(t, mins[i], min)
	}
}

func TestFind(t *testing.T) {
	v := []float64{3, 4, 5, 6, 7, math.NaN(), -1}

	f, _ := FindIndex(v, 5)
	assert.Equal(t, 2, f)
	f, _ = FindIndex(v, 999)
	assert.Equal(t, -1, f)

	c := Contains(v, 5)
	assert.Equal(t, true, c)
	c = Contains(v, 999)
	assert.Equal(t, false, c)
	c = Contains(v, math.NaN())
	assert.Equal(t, false, c)
}

func TestSliceOf(t *testing.T) {
	v := SliceOf(9, 9)
	assert.Equal(t, 9, len(v))
	v = Insert(v, 5, 1.1)
	assert.Equal(t, 10, len(v))
	assert.Equal(t, 1.1, v[5])
	v = RemoveFrom(v, 2)
	assert.Equal(t, 9, len(v))
	assert.Equal(t, 1.1, v[4])
}

func TestCumsum(t *testing.T) {
	v := SliceOf(2, 10)
	r := Cumsum(v)
	expected := []float64{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}
	assert.Equal(t, expected, r)
	r = MulSlices(v, r)
	expected = []float64{4, 8, 12, 16, 20, 24, 28, 32, 36, 40}
	assert.Equal(t, expected, r)

	v = []float64{10, math.NaN()}
	v2 := []float64{1,1}
	r = MulSlices(v, v2)
	assert.Equal(t,true,IsEqualSlice(v,r))
}
