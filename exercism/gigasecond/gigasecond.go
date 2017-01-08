package gigasecond

import (
	"math"
	"time"
)

const testVersion = 4

func AddGigasecond(t time.Time) time.Time {
	s := time.Duration(math.Pow(10, 9)) * time.Second
	p := t.Add(s)
	return p
}
