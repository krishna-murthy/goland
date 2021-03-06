package main

import "fmt"

func fibonacci() func() int {
	current, next := 0, 1
	return func() int {
		defer func() {
			tmp := current
			current, next = next, tmp+next
		}()
		return current
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
