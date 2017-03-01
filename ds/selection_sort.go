package main

import "fmt"

func main() {
	b := [10]int{64, 12, 18, 19, 17, 88, 89, 21, 56, 76}
	for i := 0; i < len(b); i++ {
		min := i
		for j := i + 1; j < len(b); j++ {
			if b[min] > b[j] {
				min = j
			}
		}
		temp := b[i]
		b[i] = b[min]
		b[min] = temp
	}
	fmt.Println(b)
}
