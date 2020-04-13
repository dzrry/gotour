package task1

import (
	"fmt"
	"math"
)

const delta = 1e-5

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("No sqrt of negative number %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x
	z0 := 0.0
	for math.Abs(z-z0) > delta {
		z0, z = z, z-(math.Pow(z, 2)-x)/(2*z)
	}
	return z, nil
}
