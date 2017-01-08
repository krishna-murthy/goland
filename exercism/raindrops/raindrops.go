package raindrops

import "fmt"

const testVersion = 2

func Convert(i int) string {
	if i%3 == 0 || i%5 == 0 || i%7 == 0 {
		var s string
		if i%3 == 0 {
			s += "Pling"
		}
		if i%5 == 0 {
			s += "Plang"
		}
		if i%7 == 0 {
			s += "Plong"
		}
		return s
	}
	return fmt.Sprintf("%d", i)
}
