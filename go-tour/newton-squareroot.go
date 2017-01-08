package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	return math.Sqrt(x)
}

func Newton(x float64) float64 {
	if x == 0 {
		return 0
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z = z - ((math.Pow(z, 2) - x) / (2 * z))
	}
	return z
}

func main() {
	for j := 0; j < 15; j++ {
		fmt.Println(j)
		fmt.Println(Sqrt(float64(j)))
		fmt.Println(Newton(float64(j)))
		fmt.Println("Diff", math.Abs(Sqrt(float64(j))-Newton(float64(j))))
	}
}
