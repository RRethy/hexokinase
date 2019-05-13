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
