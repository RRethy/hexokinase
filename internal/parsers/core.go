package parser

import (
	"bufio"
	"github.com/rrethy/hexokinase/internal/models"
	"github.com/rrethy/hexokinase/internal/output"
	"os"
)

type patParser (func(string) []*models.Colour)

// Read TODO
func Read(in *os.File) {
	scanner := bufio.NewScanner(in)
	colours := make([]*models.Colour, 0, 4)
	parsers := [](patParser){
		parseHex,
		parseRGB,
		parseRGBA,
		parseWebColours,
	}

	lnum := 0
	for scanner.Scan() {
		lnum++
		for _, parser := range parsers {
			lineColours := parser(scanner.Text())
			for _, colour := range lineColours {
				colour.Lnum = lnum
			}
			colours = append(colours, lineColours...)
		}
	}

	for _, colour := range colours {
		output.PrintColour(colour)
	}
}
