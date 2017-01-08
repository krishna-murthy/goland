package main

import "fmt"

const PI = 3.14

func main() {
	const world = "India"
	fmt.Println("Hello", world)
	fmt.Println("Happy", PI, "day")
	const Truth = true
	fmt.Println("Go Rules?", Truth)
}
