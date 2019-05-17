package main

import (
	"fmt"
)

type rgb struct {
	r, g, b int
}

type parser (func(line string) colours)

type colours []Colour

func (clrs colours) Len() int {
	return len(clrs)
}

func (clrs colours) Less(i, j int) bool {
	clr1 := clrs[i]
	clr2 := clrs[j]
	return clr1.Lnum < clr2.Lnum ||
		clr1.ColStart < clr2.ColStart ||
		clr1.ColEnd < clr2.ColEnd
}

func (clrs colours) Swap(i, j int) {
	clrs[i], clrs[j] = clrs[j], clrs[i]
}

// Colour represents a colour that was parsed and recognized as a valid colour.
type Colour struct {
	ColStart, ColEnd int
	Lnum             int
	Line             string
	Hex              string
}

func (c *Colour) String() string {
	return fmt.Sprintf("%d:%d-%d:%s:%s", c.Lnum, c.ColStart, c.ColEnd, c.Hex, c.Line)
}
