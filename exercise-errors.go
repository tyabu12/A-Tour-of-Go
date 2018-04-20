// +build ignore

package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func Sqrt(x float64) (float64, error) {
	if x < 0. {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.
	for math.Abs(z*z-x) > 0.0000000000001 {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func (e ErrNegativeSqrt) Error() string {
	return "cannot Sqrt negative number: " + fmt.Sprint(float64(e))
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
