package parser

import (
	"bufio"
	"fmt"
	"github.com/rrethy/hexokinase/internal/models"
	"os"
)

type patParser (func(string, int) []*models.Colour)

// Parse TODO
func Parse(in *os.File, out *os.File) {
	scanner := bufio.NewScanner(in)
	var colours []*models.Colour
	parsers := [](patParser){
		parseHex,
		parseRGB,
		parseRGBA,
	}

	lnum := 0
	for scanner.Scan() {
		lnum++
		for _, parser := range parsers {
			colours = append(colours, parser(scanner.Text(), lnum)...)
		}
	}

	for _, colour := range colours {
		fmt.Fprintf(out, "%d:%d-%d:%s\n",
			colour.Lnum, colour.ColStart, colour.ColEnd, colour.Hex)
	}
}
