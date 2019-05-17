package main

import (
	"fmt"
	"os"
)

var (
	// Out TODO
	out *os.File
)

// SetOut TODO
func SetOut(f *os.File) {
	out = f
}

func printColours(clrs colours) {
	for _, colour := range clrs {
		fmt.Fprintf(out, "%d:%d-%d:%s:%s\n", colour.Lnum, colour.ColStart, colour.ColEnd, colour.Hex, colour.Line)
	}
}
