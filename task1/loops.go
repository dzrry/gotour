package task1

import "math"

const delta = 1e-5

func Sqrt(x float64) float64 {
	z := x
	z0 := 0.0
	for math.Abs(z-z0) > delta {
		z0, z = z, z-(math.Pow(z, 2)-x)/(2*z)
	}
	return z
}
