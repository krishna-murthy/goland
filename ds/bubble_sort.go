package main

import "fmt"

func main() {
	a := [10]int{64, 12, 18, 19, 17, 88, 89, 21, 56, 76}
	var swap = false
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				temp := a[j]
				a[j] = a[j+1]
				a[j+1] = temp
				swap = true
			}
		}
		if swap == false {
			break
		}
	}
	fmt.Println(a)
}
