package xydata

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

// FOR TESTING
var testx = []float64{1, 4, 2, 5, 3}
var testy = []float64{0, 0, 0, 0, 0}

var testindex = []int{0, 1, 2, 3, 4}
var testsortedindex = []int{0, 2, 4, 1, 3}
var testreverseindex = []int{3, 1, 4, 2, 0}

func TestBasic(t *testing.T) {

	xy := NewXY(testx, []float64{})
	assert.Equal(t, true, len(xy.Error) > 0) // test failing NewXY

	xy = NewXY(testx, testy)
	assert.Equal(t, 5, xy.Len()) // test length

	xy.ReplaceXY(-1, -2, 3)
	assert.Equal(t, -1.0, xy.GetX(3)) // test replace X
	assert.Equal(t, -2.0, xy.GetY(3)) // test replace Y
	xy.ReplaceXY(9, 9, 6)
	assert.Equal(t, true, len(xy.Error) > 0) // test failing indexing

}

func TestAdd(t *testing.T) {

	xy := NewXY(testx, testy)
	xy.AddXY(8.8, 9.9) // test Adding data pair

	assert.Equal(t, 8.8, xy.GetX(5)) //
	assert.Equal(t, 9.9, xy.GetY(5)) //
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, xy.Map)

}
func TestRemove(t *testing.T) {

	xy := NewXY(testx, testy)
	xy.RemoveXY(1) // test Remove data pair

	assert.Equal(t, testx[2], xy.GetX(1))
	assert.Equal(t, testy[2], xy.GetY(1))
	assert.Equal(t, []int{0, 1, 2, 3}, xy.Map)

}

func TestInsert(t *testing.T) {

	xy := NewXY(testx, testy)
	xy.InsertXY(6.6, 7.7, 1) // test Insert data pair

	assert.Equal(t, testx[1], xy.GetX(2))
	assert.Equal(t, testy[1], xy.GetY(2))
	assert.Equal(t, 6.6, xy.GetX(1))
	assert.Equal(t, 7.7, xy.GetY(1))
	assert.Equal(t, testx[0], xy.GetX(0))
	assert.Equal(t, testy[0], xy.GetY(0))
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5}, xy.Map)

}

func TestFindIndex(t *testing.T) {

	xy := NewXY(testx, testy)
	assert.Equal(t, 0, xy.FindIndex(math.NaN(), testy[1]))
	assert.Equal(t, 2, xy.FindIndex(testx[2], math.NaN()))
	assert.Equal(t, -1, xy.FindIndex(math.NaN(), math.NaN()))
	assert.Equal(t, -1, xy.FindIndex(999, 999))
	assert.Equal(t, 4, xy.FindIndex(testx[4], testy[4]))
}
