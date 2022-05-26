package main

import (
	"fmt"
	"math"
)

func MySqrt(x float64) float64 {
	z := x
	epsilon := 0.000000000000001
	for {
		next := z - (z*z-x)/(2*z)
		if z == next || math.Abs(z-next) < epsilon {
			return z
		}
		fmt.Println("current z :", z)
		z -= (z*z - x) / (2 * z)
	}
}

func main() {
	fmt.Println("  my sqrt :", MySqrt(2))
	fmt.Println("math.Sqrt :", math.Sqrt(2))

}
