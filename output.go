package main

import (
	"fmt"
	"os"
)

// FMT is the colour format which will be printed to stdout
type FMT int

const (
	// ExtendedFmt will print the format "filename:lnum:colstart-colend:hex:line" (separator is ; in Windows)
	ExtendedFmt FMT = iota
	// ShortFmt will print the format "filename:lnum:colstart-colend:hex" (separator is ; in Windows)
	ShortFmt

	sep         = string(os.PathListSeparator)
	extendedFmt = "%s" + sep + "%d" + sep + "%d-%d" + sep + "%s" + sep + "%s\n"
	shortFmt    = "%s" + sep + "%d" + sep + "%d-%d" + sep + "%s\n"
)

var outputFmt = ExtendedFmt

// SetOutputFmt sets the format to print the colour
// fmt can be hexokinase.ExtendedFmt or hexokinase.ShortFmt
func SetOutputFmt(fmt FMT) {
	outputFmt = fmt
}

// PrintColours prints clrs to out based on the formatting specificed by
// SetOutputFmt.
func PrintColours(clrs colours, out *os.File, reverse bool) {
	if reverse {
		for i := len(clrs) - 1; i >= 0; i-- {
			printColour(clrs[i], out)
		}
	} else {
		for _, colour := range clrs {
			printColour(colour, out)
		}
	}
}

func printColour(colour *Colour, out *os.File) {
	if outputFmt == ExtendedFmt {
		fmt.Fprintf(out, extendedFmt, colour.Tag, colour.Lnum, colour.ColStart, colour.ColEnd, colour.Hex, colour.Line)
	} else {
		fmt.Fprintf(out, shortFmt, colour.Tag, colour.Lnum, colour.ColStart, colour.ColEnd, colour.Hex)
	}
}
