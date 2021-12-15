package xydata

import "github.com/Akkurate/utils/numf"

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
