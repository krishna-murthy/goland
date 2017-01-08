package hamming

import (
	"fmt"
	"strings"
)

type MyError struct {
	What string
}

func (e MyError) Error() string {
	return fmt.Sprintf("error: %s\n", e.What)
}

const testVersion = 5

func Distance(a, b string) (int, error) {
	var err error
	c := 0
	a1 := strings.Split(a, "")
	b1 := strings.Split(b, "")
	if len(a1) == len(b1) {
		for i := 0; i < len(a1); i++ {
			if a1[i] != b1[i] {
				c += 1
			}
		}
	} else {
		err = MyError{"length dint match"}
		c = -1
	}
	return c, err
}
