package numf

import (
	"fmt"
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
	DropNanInplace(&vals)
	assert.Equal(t, []float64{0, 1, -1}, vals)

	vals = []float64{1,2,3,math.NaN(),5,6,math.NaN()}
	ReplaceNansInplace(&vals,0)
	assert.Equal(t, []float64{1,2,3,0,5,6,0}, vals)

}

func TestFillnan(t *testing.T) {
	nan := math.NaN()
	v := []float64{nan, 0, 1, 2, nan, 4, nan, nan}
	rprev := []float64{nan, 0, 1, 2, 2, 4, 4, 4}
	rlin := []float64{nan, 0, 1, 2, 3, 4, 4, 4}

	resp := FillNan(v, "previous", false, 0)
	assert.Equal(t, true, IsEqualSlice(rprev, resp))

	resp = FillNan(v, "linear", false, 0)
	assert.Equal(t, true, IsEqualSlice(rlin, resp), fmt.Sprintf("resp %v != %v", resp, rlin))

	resp = FillNan(v, "linear", true, 0)
	assert.Equal(t, true, IsEqualSlice(ReplaceNans(rlin, 0), resp), fmt.Sprintf("resp %v != %v", resp, ReplaceNans(rlin, 0)))

	long := NanSlice(20)
	long[5] = 1
	long[16] = 2

	resp = FillNan(long, "previous", false, 5)
	expected := []float64{nan, nan, nan, nan, nan, 1, 1, 1, 1, 1, nan, nan, nan, nan, nan, nan, 2, 2, 2, 2}
	assert.Equal(t, true, IsEqualSlice(expected, resp), fmt.Sprintf("resp %v != %v", resp, expected))

	resp = FillNan(long, "previous", true, 5)
	expected = []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, nan, nan, nan, nan, nan, nan, 2, 2, 2, 2}
	assert.Equal(t, true, IsEqualSlice(expected, resp), fmt.Sprintf("resp %v != %v", resp, expected))

	resp = FillNan(long, "previous", true, 2)
	expected = []float64{1, 1, nan, nan, nan, 1, 1, nan, nan, nan, nan, nan, nan, nan, nan, nan, 2, 2, nan, nan}
	assert.Equal(t, true, IsEqualSlice(expected, resp), fmt.Sprintf("resp %v != %v", resp, expected))
}

func TestFillnanInplace(t *testing.T) {

	nan := math.NaN()
	testset := []float64{nan, 0, 1, 2, nan, 4, nan, nan}
	rprev := []float64{nan, 0, 1, 2, 2, 4, 4, 4}
	rlin := []float64{nan, 0, 1, 2, 3, 4, 4, 4}

	FillNanInplace(&testset, "previous", false, 0)
	assert.Equal(t, true, IsEqualSlice(rprev, testset))
	testset = []float64{nan, 0, 1, 2, nan, 4, nan, nan}
	FillNanInplace(&testset, "linear", false, 0)
	assert.Equal(t, true, IsEqualSlice(rlin, testset), fmt.Sprintf("resp %v != %v", testset, rlin))
	testset = []float64{nan, 0, 1, 2, nan, 4, nan, nan}
	FillNanInplace(&testset, "linear", true, 0)
	assert.Equal(t, true, IsEqualSlice(ReplaceNans(rlin, 0), testset), fmt.Sprintf("testset %v != %v", testset, ReplaceNans(rlin, 0)))

	long := NanSlice(20)
	long[5] = 1
	long[16] = 2

	FillNanInplace(&long, "previous", false, 5)
	expected := []float64{nan, nan, nan, nan, nan, 1, 1, 1, 1, 1, nan, nan, nan, nan, nan, nan, 2, 2, 2, 2}
	assert.Equal(t, true, IsEqualSlice(expected, long), fmt.Sprintf("long %v != %v", long, expected))
	long = NanSlice(20)
	long[5] = 1
	long[16] = 2
	FillNanInplace(&long, "previous", true, 5)
	expected = []float64{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, nan, nan, nan, nan, nan, nan, 2, 2, 2, 2}
	assert.Equal(t, true, IsEqualSlice(expected, long), fmt.Sprintf("long %v != %v", long, expected))
	long = NanSlice(20)
	long[5] = 1
	long[16] = 2
	FillNanInplace(&long, "previous", true, 2)
	expected = []float64{1, 1, nan, nan, nan, 1, 1, nan, nan, nan, nan, nan, nan, nan, nan, nan, 2, 2, nan, nan}
	assert.Equal(t, true, IsEqualSlice(expected, long), fmt.Sprintf("long %v != %v", long, expected))
}
