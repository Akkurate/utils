package xydata

import (
	"github.com/Akkurate/utils/numf"
	"github.com/Akkurate/utils/numi"
)

// Length of data
func (xy *XY) Len() int {
	xy.Error = ""
	return len(xy.Map)
}

// Replaces XY pair from given index.
func (xy *XY) ReplaceXY(x, y float64, i int) {

	if xy.isBadindexing(i) {
		return
	}

	xy.data[xy.Map[i]] = xypair{x: x, y: y, i: i}
}

// Adds XY pair to the tail of data array.
func (xy *XY) AddXY(x, y float64) {

	xy.Error = ""

	newindex := len(xy.data)
	xy.data = append(xy.data, xypair{x: x, y: y, i: newindex})
	xy.Map = append(xy.Map, newindex)
}

// Inserts XY pair into data array.
func (xy *XY) InsertXY(x, y float64, i int) {

	// i index points to Map -- we need to figure out the index to real data
	// strategy is to insert to higher real index, if ambiguous
	reali := xy.Map[i] // default real index
	if i > 0 {
		if xy.Map[i-1] > xy.Map[i] {
			reali = xy.Map[i-1]
		}
	}
	xy.data = append(xy.data, xypair{})
	copy(xy.data[reali+1:], xy.data[reali:])
	xy.data[reali] = xypair{x: x, y: y, i: reali}

	// update index within pairs
	for x, p := range xy.data[reali+1:] {
		p.i = p.i + 1
		xy.data[reali+1+x] = p
	}

	// update map to reflect inserting
	for i, v := range xy.Map {
		if v >= reali {
			xy.Map[i] = v + 1
		}
	}
	xy.Map = numi.Insert(xy.Map, i, reali)

}

// Removes XY pair from given index. Note that this actually removes the data, not just the index to it.
func (xy *XY) RemoveXY(i int) {
	if xy.isBadindexing(i) {
		return
	}

	reali := xy.Map[i] // get the index to real data...

	xy.data = append(xy.data[:reali], xy.data[reali+1:]...)
	// update map to reflect removing
	for i, v := range xy.Map {
		if v > reali {
			xy.Map[i] = v - 1
		}
	}
	xy.Map = append(xy.Map[:i], xy.Map[i+1:]...)

}

// Returns X and Y from given index
func (xy *XY) GetXY(i int) (x, y float64) {
	if xy.isBadindexing(i) {
		return
	}
	return xy.data[xy.Map[i]].x, xy.data[xy.Map[i]].y
}

// Returns X from given index
func (xy *XY) GetX(i int) (x float64) {
	if xy.isBadindexing(i) {
		return
	}
	return xy.data[xy.Map[i]].x
}

// Returns Y from given index
func (xy *XY) GetY(i int) (y float64) {
	if xy.isBadindexing(i) {
		return
	}
	return xy.data[xy.Map[i]].y
}

// Finds the index of given X and/or Y value(s). Index is gone through in ascending order and the first match is returned.
// Use math.NaN() to disable matching. Function returns -1 is not match is found.
func (xy *XY) FindIndex(x, y float64) int {
	xy.Error = ""
	matchboth := numf.IsValid(x) && numf.IsValid(y)
	for i, v := range xy.Map {
		p := xy.data[v]
		switch matchboth {
		case true:
			if p.x == x && p.y == y {
				return i
			}
		case false:
			if p.x == x || p.y == y {
				return i
			}
		}
	}
	return -1
}
