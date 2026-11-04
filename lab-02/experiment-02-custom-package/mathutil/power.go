package mathutil

// Power calculates base raised to the power of exponent
func Power(base, exponent int) int {
	if exponent < 0 {
		return 0
	}
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
