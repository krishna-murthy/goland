package main

import "fmt"

func main() {
	var i interface{}

	i = "hello"
	t := i.(string)
	fmt.Println(t)

	t, ok := i.(string)
	fmt.Println(t, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) //panic
	fmt.Println(f)
}
