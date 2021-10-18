package logic

// Xor performs exclusive OR operation
func Xor(a, b bool) bool {
	return a != b
}

// Xnor performs exclusive NOR operation
func Xnor(a, b bool) bool {
	return !((a || b) && (!a || !b))
}
