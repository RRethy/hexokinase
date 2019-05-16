package main

import (
	"fmt"
)

// Colour TODO
type Colour struct {
	ColStart, ColEnd int
	Lnum             int
	Hex              string
}

func (c *Colour) String() string {
	return fmt.Sprintf("%d:%d-%d:%s", c.Lnum, c.ColStart, c.ColEnd, c.Hex)
}
