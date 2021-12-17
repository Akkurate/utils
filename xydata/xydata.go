/* XYDATA

A data structure for a set of (x,y) data pairs with basic functionality access and manipulate data, as well sort, filter and apply functions to data.
Data accessing is done by indexing to data array via Map or "active data index". For example when sorting, the actual data is not sorted but the Map
is re-ordered in a way that the data appears sorted.
*/
package xydata

import "github.com/Akkurate/utils/numi"

// single xy data pair
type xypair struct {
	x float64 // X value
	y float64 // Y value
	i int     // index; needed for sorting
}

// xy data
type XY struct {
	data  []xypair // XY-data array -- data is accessed with provided functions and methods, using Map when applicable
	Map   []int    // Maps the active data index to real data index -- used to access desired data in desired order outside XY
	Error string   // Error string -- cleared when any new XY function call is made
}

// Creates new XY object
func NewXY(x, y []float64) *XY {

	if len(x) != len(y) {
		xy := &XY{Error: "Mismatch in input data length."}
		return xy
	}
	if len(x) == 0 || len(y) == 0 {
		xy := &XY{Error: "No input data."}
		return xy
	}

	xy := &XY{
		data:  make([]xypair, len(x)),
		Error: "",
	}

	// fill array with data pairs & index
	for i := 0; i < len(x); i++ {
		p := xypair{x: x[i], y: y[i], i: i}
		xy.data[i] = p
	}
	xy.Map = numi.NumRange(0, len(xy.data)-1, 1) // create Map; initially a copy of real data index
	return xy
}

// Resets Map to reflect real data index
func (xy *XY) ResetMap() {
	xy.Map = numi.NumRange(0, len(xy.data)-1, 1)
}

// Error checks for indexing. Note that indexing error do not cause any panicking etc.; Error -string need to be checked outside the function to see any issues.
func (xy *XY) isBadindexing(i int) bool {
	if i < 0 || i >= len(xy.Map) {
		xy.Error = "Out of index."
		return true
	}
	xy.Error = ""
	return false
}


