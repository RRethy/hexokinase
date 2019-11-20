package main

import (
	"fmt"
)

type rgb struct {
	r, g, b int
}

type parser (func(line string) colours)

type colours []*Colour

func (clrs colours) Len() int {
	return len(clrs)
}

func (clrs colours) Less(i, j int) bool {
	clr1 := clrs[i]
	clr2 := clrs[j]
	if clr1.Lnum < clr2.Lnum ||
		clr1.Lnum == clr2.Lnum && clr1.ColStart <= clr2.ColStart ||
		clr1.Lnum == clr2.Lnum && clr1.ColStart == clr2.ColStart && clr1.ColEnd <= clr2.ColEnd {
		return true
	}
	return false
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
	Tag              string
}

func (c *Colour) String() string {
	return fmt.Sprintf("%s:%d:%d-%d:%s:%s", c.Tag, c.Lnum, c.ColStart, c.ColEnd, c.Hex, c.Line)
}
