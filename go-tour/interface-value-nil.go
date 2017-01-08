package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func describe(i I) {
	fmt.Printf("%v %T\n", i, i)
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t)
}

func main() {
	var i I
	var t *T
	i = t
	describe(i)
	i.M()
	i = &T{"hello"}
	describe(i)
	i.M()
}
