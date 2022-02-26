package misc

import (
	"sort"

	"github.com/Akkurate/utils/numf"
)
// Inserts a float into a slice, sorts the slice and returns the index of first appearance of given float after sorting
func InsertSortFindIndex(s []float64, v float64) int {
	l := make([]float64, len(s))
	copy(l, s)
	l = append(l, v)
	sort.Float64s(l)
	i, _ := numf.FindIndex(l, v)
	return i
}
