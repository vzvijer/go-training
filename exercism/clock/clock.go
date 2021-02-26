package clock

import "fmt"

type Clock struct {
	h int
	m int
}

func New(h, m int) Clock {
	h = (h + (m / 60)) % 24
	m = m % 60
	if m < 0 {
		m = m + 60
		h--
	}
	if h < 0 {
		h = h + 24
	}
	return Clock{h, m}
}

func (c Clock) Add(m int) Clock {
	return New(c.h, c.m+m)
}

func (c Clock) Subtract(m int) Clock {
	return New(c.h, c.m-m)
}

func (c Clock) String() string {
	hs := fmt.Sprintf("%d", c.h)
	if c.h < 10 {
		hs = "0" + hs
	}
	ms := fmt.Sprintf("%d", c.m)
	if c.m < 10 {
		ms = "0" + ms
	}
	return hs + ":" + ms
}
