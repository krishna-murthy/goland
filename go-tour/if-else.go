package main

import (
	"fmt"
	"math"
)

func limit(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	fmt.Println(limit(3, 2, 10))
	fmt.Println(limit(3, 3, 20))
}
