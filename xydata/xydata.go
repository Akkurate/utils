package xydata

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


