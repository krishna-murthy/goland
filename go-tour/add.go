package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	//    fmt.Println(add(42,43))
	// Another way for add
	sum := func(a, b int) int { return a + b }(3, 4)
	fmt.Println(sum)
}
