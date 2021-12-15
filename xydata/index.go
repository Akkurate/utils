package xydata

import (
	"sort"

	"github.com/Akkurate/utils/numi"
)

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

// Filters all X data with given function and returns the index of filtered values
func (xy *XY) FilterX(f func(x float64) bool) []int {
	var index []int
	for i, p := range xy.data {
		if f(p.x) {
			index = append(index, i)
		}
	}
	return index
}

// Filters all Y data with given function and returns the index of filtered values
func (xy *XY) FilterY(f func(y float64) bool) []int {
	var index []int
	for i, p := range xy.data {
		if f(p.y) {
			index = append(index, i)
		}
	}
	return index
}
