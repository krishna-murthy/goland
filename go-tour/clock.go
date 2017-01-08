// Clock stub file

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package main

import (
    "fmt"
    "time"
)
// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.

type Clock struct {
    t time.Time
}

func New(hour, minute int) Clock {
    var r Clock
    t := time.Date(1988, time.January, 1, hour, 0, 0, 0, time.UTC)
    m := time.Duration(minute) * time.Minute
    s := t.Add(m)
    r.t = s
    return r
}

func (c Clock) String() string {
    t := c.t.Minute()
    s := c.t.Hour()
    return fmt.Sprintf("%02d:%02d\n", s, t)
}

func (c Clock) Add(minutes int) Clock {
    t := c.t
    m := time.Duration(minutes) * time.Minute
    s := t.Add(m)
    c.t = s
    return c
}

func main() {
    var g Clock
    g = New(7, 32)
    t := g.t
    fmt.Println(t.Day())
    var c Clock
    c = New(-12, -268)
    t = c.t
    fmt.Println(t.Day())
    fmt.Println(g == c)
    fmt.Println(g.String())
}
