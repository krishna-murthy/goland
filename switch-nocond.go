package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 18:
		fmt.Println("Good afternoon!")
	case t.Hour() < 20:
		fmt.Println("Good evening!")
	default:
		fmt.Println("Good night! Go sleep!")
	}
}
