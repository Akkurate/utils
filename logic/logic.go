/* Logic operations.
 */
package logic

// Performs exclusive OR operation.
func Xor(a, b bool) bool {
	return a != b
}

// Performs exclusive NOR operation.
func Xnor(a, b bool) bool {
	return !((a || b) && (!a || !b))
}
