package main

import "fmt"

func main() {
	do(21)
	do("hello")
	do(false)
}

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Its an int with value: %d\n", v)
	case string:
		fmt.Printf("Its a string saying: %v\n", v)
	default:
		fmt.Printf("Not sure what this is : %T of value %v\n", v, v)
	}
}
