package sqrt

func Sqrt(x float64) float64 {
	z := float64(1)
	var val float64
	for z != val {
		val = z
		z -= (z*z - x) / (2 * z)
	}
	return z
}
