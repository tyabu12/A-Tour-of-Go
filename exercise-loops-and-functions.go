// +build ignore

package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	z, cnt := 1., 0
	for math.Abs(z*z-x) > 0.0000000000001 {
		z -= (z*z - x) / (2 * z)
		cnt++
	}
	return z, cnt
}

func main() {
	z, cnt := Sqrt(2)
	fmt.Println(z, z-math.Sqrt(2), cnt)
}
