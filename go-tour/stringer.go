package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}

func main() {
	p := Person{"Krishna", 28}
	q := Person{"Surya", 24}
	fmt.Println(p, q)
}
