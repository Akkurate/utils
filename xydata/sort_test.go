package xydata

import (
	"testing"

	"github.com/stretchr/testify/assert"
)


func TestSort(t *testing.T) {

	xy := NewXY(testx, testy)
	xy.SortX(true) // sort by X ascending...
	assert.Equal(t, testsortedindex, xy.Map)
	assert.Equal(t, 5, len(xy.Map))

	xy.ResetMap() // reset index
	assert.Equal(t, testindex, xy.Map)
	assert.Equal(t, 5, len(xy.Map))

	xy = NewXY(testy, testx)
	xy.SortY(false) // sort by Y descending
	assert.Equal(t, testreverseindex, xy.Map)
	assert.Equal(t, 5, len(xy.Map))

	xy.SortI(true) // sort by XY pair indexes ascending
	assert.Equal(t, testindex, xy.Map)
	assert.Equal(t, 5, len(xy.Map))
}
