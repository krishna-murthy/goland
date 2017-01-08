package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 0}
		v3 = Vertex{}
		p  = &Vertex{1, 2}
	)
	fmt.Println(v1, v2, v3, p)
}
