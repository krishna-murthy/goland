package main

import "fmt"

var pow = []int{1, 4, 9, 16, 25, 36}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
