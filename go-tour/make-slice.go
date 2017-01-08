package main

import "fmt"

func main() {
	s := make([]int, 5)
	printSlice("s", s)

	p := make([]int, 0, 5)
	printSlice("p", p)

	q := p[:2]
	printSlice("q", q)

	v := q[2:5]
	printSlice("v", v)
}

func printSlice(a string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", a, len(x), cap(x), x)
}
