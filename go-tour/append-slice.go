package main

import "fmt"

func main() {
	var s = []int{1, 2}
	fmt.Printf("%v\n", s)

	s = append(s, 2)
	fmt.Printf("%v\n", s)

	s = append(s, 3, 4, 5)
	fmt.Printf("%v\n", s)
}
