package main

import "fmt"

func main() {
	names := [4]string{
		"James",
		"House",
		"Cameron",
		"Chase",
	}
	var a []string = names[0:2]
	var b []string = names[1:3]
	fmt.Println(a, b)

	a[1] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
