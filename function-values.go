package main

import (
    "fmt"
    "math"
)

func compute(fn func(float64, float64)float64) float64 {
    return fn(3, 4)
}

func main() {
    hyplot := func(x, y float64) float64 {
        return math.Sqrt( x*x + y*y )
    }

    fmt.Println(hyplot(5, 12))

    fmt.Println(compute(hyplot))
    fmt.Println(compute(math.Pow))
}
