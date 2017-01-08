package clock

import (
	"fmt"
	"time"
)

const testVersion = 4

type Clock struct {
	t time.Time
}

func New(hour, minute int) Clock {
	var r Clock
	hour %= 24
	minute %= 1440
	if hour < 0 && minute < 0 {
		for minute < -60 {
			minute += 60
			hour -= 1
		}
		hour -= 1
		hour = (hour % 24) + 24
		minute = 60 + minute
	} else if hour < 0 {
		hour = 24 + hour
	}
	for minute < -60 {
		minute += 1440
	}
	t := time.Date(0, time.January, 1, hour, minute, 0, 0, time.UTC)

	r.t = t
	return r
}

func (c Clock) String() string {
	t := c.t.Minute()
	s := c.t.Hour()
	return fmt.Sprintf("%02d:%02d", s, t)
}

func (c Clock) Add(minutes int) Clock {
	t := c.t
	m := time.Duration(minutes) * time.Minute
	s := t.Add(m)
	c.t = s
	return c
}
