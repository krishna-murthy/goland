package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func scaleVertex(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	scaleVertex(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(2)
	scaleVertex(p, 10)

	fmt.Println(v, p)
}
