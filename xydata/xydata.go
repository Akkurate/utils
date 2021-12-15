package xydata

import (
	"sort"

	"github.com/Akkurate/utils/numf"
	"github.com/Akkurate/utils/numi"
)

// single xy data pair
type xypair struct {
	x float64 // X value
	y float64 // Y value
	i int     // index; needed for sorting
}

// xy data
type XY struct {
	data  []xypair // XY-data array -- not to be accessed directly
	Error string   // Error string -- cleared when new function call is made to XYdata object
}

// Creates new XY object
func NewXY(x, y []float64) *XY {

	if len(x) != len(y) {
		xyd := &XY{Error: "Mismatch in input data length"}
		return xyd
	}

	xyd := &XY{
		data:  make([]xypair, len(x)),
		Error: "",
	}
	// fill array with data pairs & index
	for i := 0; i < len(x); i++ {
		p := xypair{x: x[i], y: y[i], i: i}
		xyd.data[i] = p
	}

	return xyd
}

// Length of data
func (xy *XY) Len() int {
	xy.Error = ""
	return len(xy.data)
}

// Replaces XY pair from given index.
func (xy *XY) ReplaceXY(x, y float64, i int) {
	xy.Error = ""
	if i >= len(xy.data) {
		xy.Error = "Out of index"
		return
	}

	xy.data[i] = xypair{x: x, y: y, i: i}

}

// Adds XY pair to the tail of data array.
func (xy *XY) AddXY(x, y float64) {
	xy.Error = ""
	newindex := len(xy.data)
	xy.data = append(xy.data, xypair{x: x, y: y, i: newindex})
}

// Inserts XY pair into data array.
func (xy *XY) InsertXY(x, y float64, i int) {
	xy.Error = ""

	xy.data = append(xy.data, xypair{})
	copy(xy.data[i+1:], xy.data[i:])
	xy.data[i] = xypair{x: x, y: y, i: i}

	// update index within pairs
	for x, p := range xy.data[i+1:] {
		p.i = p.i + 1
		xy.data[i+1+x] = p
	}

}

// Removes XY pair from given index.
func (xy *XY) RemoveXY(i int) {
	xy.Error = ""
	if i >= len(xy.data) {
		xy.Error = "Out of index"
		return
	}
	xy.data = append(xy.data[:i], xy.data[i+1:]...)

}

// Sorts the X values as ascending or descending and returns a sorted index.
func (xy *XY) SortX(ascending bool) []int {
	xy.Error = ""
	sortedIndex := make([]int, len(xy.data))

	pairs := make([]xypair, len(xy.data))
	copy(pairs, xy.data)

	if ascending {
		sort.Slice(pairs, func(i, j int) bool { return pairs[i].x < pairs[j].x })
	} else {
		sort.Slice(pairs, func(i, j int) bool { return pairs[i].x > pairs[j].x })
	}

	for x, p := range pairs {
		sortedIndex[x] = p.i
	}
	return sortedIndex
}

// Sorts the Y values as ascending or descending and returns a sorted index.
func (xy *XY) SortY(ascending bool) []int {
	xy.Error = ""
	sortedIndex := make([]int, len(xy.data))

	pairs := make([]xypair, len(xy.data))
	copy(pairs, xy.data)

	if ascending {
		sort.Slice(pairs, func(i, j int) bool { return pairs[i].y < pairs[j].y })
	} else {
		sort.Slice(pairs, func(i, j int) bool { return pairs[i].y > pairs[j].y })
	}

	for x, p := range pairs {
		sortedIndex[x] = p.i
	}
	return sortedIndex
}

// Sorts the index ascending or descending and returns a sorted index.
func (xy *XY) SortI(ascending bool) []int {
	xy.Error = ""
	index := numi.NumRange(0, len(xy.data)-1, 1)
	if ascending {
		return index
	}
	sort.Slice(index, func(i, j int) bool { return index[j] < index[i] })
	return index
}

// Returns X and Y from given index
func (xy *XY) GetXY(i int) (x, y float64) {
	xy.Error = ""
	return xy.data[i].x, xy.data[i].y
}

// Returns X from given index
func (xy *XY) GetX(i int) (x float64) {
	xy.Error = ""
	return xy.data[i].x
}

// Returns Y from given index
func (xy *XY) GetY(i int) (y float64) {
	xy.Error = ""
	return xy.data[i].y
}

// Finds the index of given X and/or Y value(s). Index is gone through in ascending order and the first match is returned.
// Use math.NaN() to disable matching.
func (xy *XY) GetI(x, y float64) int {
	xy.Error = ""
	matchboth := numf.IsValid(x) && numf.IsValid(y)
	for _, p := range xy.data {

		switch matchboth {
		case true:
			if p.x == x && p.y == y {
				return p.i
			}
		case false:
			if p.x == x || p.y == y {
				return p.i
			}
		}
	}
	return -1
}

// Returns a view of X data in ascending index order.
func (xy *XY) ViewX() *[]float64 {
	xy.Error = ""
	view := make([]float64, len(xy.data))
	for i, p := range xy.data {
		view[i] = p.x
	}
	return &view
}

// Returns a view of Y data in ascending index order.
func (xy *XY) ViewY() *[]float64 {
	xy.Error = ""
	view := make([]float64, len(xy.data))
	for i, p := range xy.data {
		view[i] = p.y
	}
	return &view
}
