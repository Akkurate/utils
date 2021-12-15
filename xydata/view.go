package xydata

// Returns a view of X data with given index. If no index is given, all X data is shown in ascending index order.
func (xy *XY) ViewX(index []int) *[]float64 {
	xy.Error = ""
	if len(index) == 0 {
		view := make([]float64, len(xy.data))
		for i, p := range xy.data {
			view[i] = p.x
		}
		return &view
	}
	view := make([]float64, len(index))
	for i, v := range index {
		view[i] = xy.data[v].x
	}
	return &view
}

// Returns a view of Y data with given index. If no index is given, all Y data is shown in ascending index order.
func (xy *XY) ViewY(index []int) *[]float64 {
	xy.Error = ""
	if len(index) == 0 {
		view := make([]float64, len(xy.data))
		for i, p := range xy.data {
			view[i] = p.y
		}
		return &view
	}
	view := make([]float64, len(index))
	for i, v := range index {
		view[i] = xy.data[v].y
	}
	return &view
}