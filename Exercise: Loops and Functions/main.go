package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) (float64, int) {
	z := float64(1)
	i := 0
	for ; i <= 10; i++ {
		prev := z
		z -= (z*z - x) / (2 * z)

		if math.Abs(prev-z) < 0.00001 {
			break
		}
	}
	return z, i
}

func main() {
	fmt.Println(Sqrt(2.5))
	fmt.Println(math.Sqrt(2.5))
}