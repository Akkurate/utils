package xydata

import (
	"sort"
)

// Sorts the active X values ascending or descending.
func (xy *XY) SortX(ascending bool) {
	xy.Error = ""

	// copy the active data
	pairs := make([]xypair, len(xy.Map))
	for i, v := range xy.Map {
		pairs[i] = xy.data[v]
	}
	// sort the active data
	if ascending {
		sort.Slice(pairs, func(i, j int) bool { return pairs[i].x < pairs[j].x })
	} else {
		sort.Slice(pairs, func(i, j int) bool { return pairs[i].x > pairs[j].x })
	}
	// rewrite Map
	for i, p := range pairs {
		xy.Map[i] = p.i
	}
}

// Sorts the active Y values ascending or descending.
func (xy *XY) SortY(ascending bool) {
	xy.Error = ""

	pairs := make([]xypair, len(xy.Map))
	for i, v := range xy.Map {
		pairs[i] = xy.data[v]
	}

	if ascending {
		sort.Slice(pairs, func(i, j int) bool { return pairs[i].y < pairs[j].y })
	} else {
		sort.Slice(pairs, func(i, j int) bool { return pairs[i].y > pairs[j].y })
	}

	for i, p := range pairs {
		xy.Map[i] = p.i
	}

}

// Sorts the active index ascending or descending.
func (xy *XY) SortI(ascending bool) {
	xy.Error = ""

	index := xy.Map
	if ascending {
		sort.Slice(index, func(i, j int) bool { return index[j] > index[i] })
	} else {
		sort.Slice(index, func(i, j int) bool { return index[j] < index[i] })
	}
	xy.Map = index
}

