package main

import "fmt"

func main() {
	i, j := 21, 441

	p := &i
	fmt.Println(*p)
	*p = 20
	fmt.Println(i)

	p = &j
	fmt.Println(j)
	*p = *p / 21
	fmt.Println(j)
}
