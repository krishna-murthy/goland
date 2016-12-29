package main

import "fmt"

func main() {
	i := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(i)

	s := []string{"Why", "Dr.", "House"}
	fmt.Println(s)

	p := []struct {
		X int
		Y bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, false},
	}
	fmt.Println(p)
}
