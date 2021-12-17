package xydata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestView(t *testing.T) {
	xy := NewXY(testx, testy)

	viewx, viewy := xy.ViewXY()

	for i, v := range *viewx {
		assert.Equal(t, testx[i], v)
	}
	for i, v := range *viewy {
		assert.Equal(t, testy[i], v)
	}
}

func TestFilter(t *testing.T) {
	xy := NewXY(testx, testy)

	f := func(x float64) bool { return x != 0 }

	xy.FilterX(f)
	assert.Equal(t, testindex, xy.Map)
	xy.FilterY(f)
	assert.Equal(t, 0, xy.Len())
	xy.FilterX(f)
	assert.Equal(t, 0, xy.Len())
	xy.ResetMap()
	assert.Equal(t, testindex, xy.Map)
}

func TestApply(t *testing.T) {
	xy := NewXY(testx, testy)

	f := func(x, y float64) (float64, float64) { return x + 1, y - 1 }

	xy.Apply(f)
	assert.Equal(t, testx[0]+1, xy.GetX(0))
	assert.Equal(t, testy[2]-1, xy.GetY(2))
}

func TestCombo(t *testing.T) {
	xy := NewXY(testx, testy)

	n := 0.5
	f := func(x, y float64) (float64, float64) { return x + n, y - n }
	f2 := func(x float64) bool { return x > 2 }
	f3 := func(x float64) bool { return x > 2+n }
	xy.Apply(f)
	assert.Equal(t, testindex, xy.Map)
	xy.FilterX(f2)
	assert.Equal(t, []int{1, 2, 3, 4}, xy.Map)
	xy.SortX(false)
	assert.Equal(t, []int{3, 1, 4, 2}, xy.Map)
	xy.FilterX(f3)
	assert.Equal(t, []int{3, 1, 4}, xy.Map)
	assert.Equal(t, 5.5, xy.GetX(0))
	assert.Equal(t, -0.5, xy.GetY(0))
	xy.FilterY(f3)
	assert.Equal(t,0,xy.Len())
}
