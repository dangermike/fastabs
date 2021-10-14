package fastabs

func Abs(v int) int {
	temp := v >> 63 // just the sign bit
	v ^= temp       // toggle the bits if value is negative
	v += temp & 1
	return v
}
