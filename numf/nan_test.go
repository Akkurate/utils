package numf

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNans(t *testing.T) {
	vals := []float64{0, 1, -1, math.NaN(), math.Inf(1), math.Inf(-1)}
	resp := []bool{true, true, true, false, false, false}
	resp2 := []float64{0, 1, -1, 2, 2, 2}
	for i, v := range vals {
		assert.Equal(t, resp[i], IsValid(v))
		assert.Equal(t, resp2[i], ReplaceNan(v, 2))
	}
	replace := ReplaceNans(vals, 2)
	assert.Equal(t, resp2, replace)

	nanslice := NanSlice(10)
	assert.Equal(t, 10, len(nanslice))
	nanslice = DropNan(nanslice)
	assert.Equal(t, 0, len(nanslice))
	nanslice = DropNan(nanslice)
	assert.Equal(t, 0, len(nanslice))
	nanslice = DropNan(vals)
	assert.Equal(t, []float64{0, 1, -1}, nanslice)

}

func TestFillnan(t *testing.T) {
	v := []float64{math.NaN(), 0, 1, 2, math.NaN(), 4, math.NaN()}
	rprev := []float64{math.NaN(), 0, 1, 2, 2, 4, 4}
	rlin := []float64{math.NaN(), 0, 1, 2, 3, 4, 4}

	resp := FillNan(v, "previous", false, 0)
	assert.Equal(t, true, IsEqualSlice(rprev, resp))
	resp = FillNan(v, "linear", false, 0)
	assert.Equal(t, true, IsEqualSlice(rlin, resp))
	resp = FillNan(v, "linear", true, 0)
	assert.Equal(t, true, IsEqualSlice(ReplaceNans(rlin, 0), resp))

	long := NanSlice(20)
	long[5] = 1
	long[16] = 2
	resp = FillNan(long,"previous",true,5)
	n:=math.NaN()
	expected := []float64{n,n,n,n,n,1,1,1,1,1,n,n,n,n,n,n,2,2,2,2}
	assert.Equal(t, true, IsEqualSlice(expected,resp))
}
