package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	s := 1e-15

	for {
		currentZ := z - (z*z-x)/(2*z)

		if math.Abs(currentZ-z) < s {
			return currentZ
		}

		z = currentZ
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
