package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Date(0, time.January, 1, 03, 04, 0, 0, time.UTC)
	fmt.Println(t)
	m := time.Duration(-19) * time.Minute
	fmt.Println(m)
	s := t.Add(m)
	fmt.Println(s)
	fmt.Println(s.Minute())
	fmt.Println(t.Format("3:04"))
	fmt.Println(t.Format(time.RFC3339))
}
