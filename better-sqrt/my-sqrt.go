package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := x
	epsilon := 0.000000000000001
	for {
		next := z - (z*z-x)/(2*z)
		if z == next || math.Abs(z-next) < epsilon {
			return z, nil
		}
		fmt.Println("current z :", z)
		z -= (z*z - x) / (2 * z)
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
