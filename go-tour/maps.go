package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Google": {13.121231, 14.121211},
}

func main() {
	m["Bell Labs"] = Vertex{
		12.131231, 12.412311,
	}
	fmt.Println(m)
}
