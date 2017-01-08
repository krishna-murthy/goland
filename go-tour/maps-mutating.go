package main

import "fmt"

func main() {
	var m = make(map[string]int)

	m["Answer"] = 42
	fmt.Println(m["Answer"])

	delete(m, "Answer")
	fmt.Println(m["Answer"])

	elem, ok := m["Answer"]
	fmt.Println(elem, ok)
}
