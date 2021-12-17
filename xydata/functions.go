package xydata

// Returns a view of active X data.
func (xy *XY) ViewX() *[]float64 {

	xy.Error = ""
	view := make([]float64, len(xy.Map))
	for i, v := range xy.Map {
		view[i] = xy.data[v].x
	}
	return &view
}

// Returns a view of active Y data.
func (xy *XY) ViewY() *[]float64 {

	xy.Error = ""
	view := make([]float64, len(xy.Map))
	for i, v := range xy.Map {
		view[i] = xy.data[v].y
	}
	return &view
}

// Returns views to active X and Y data.
func (xy *XY) ViewXY() (x *[]float64, y *[]float64) {
	return xy.ViewX(), xy.ViewY()
}

// Filters all X data with given function.
func (xy *XY) FilterX(f func(x float64) bool) {
	var index []int
	for _, v := range xy.Map {
		p := xy.data[v]
		if f(p.x) {
			index = append(index, p.i)
		}
	}
	xy.Map = index
}

// Filters all Y data with given function.
func (xy *XY) FilterY(f func(y float64) bool) {
	var index []int
	for _, v := range xy.Map {
		p := xy.data[v]
		if f(p.y) {
			index = append(index, p.i)
		}
	}
	xy.Map = index
}

// Apply function to active data.
func (xy *XY) Apply(f func(x, y float64) (float64, float64)) {
	for _, v := range xy.Map {
		xy.data[v].x, xy.data[v].y = f(xy.data[v].x, xy.data[v].y)
	}
}
