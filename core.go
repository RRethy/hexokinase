package main

import (
	"bufio"
	"os"
)

type patParser (func(string) []*Colour)

// Read TODO
func Read(in *os.File) {
	scanner := bufio.NewScanner(in)
	colours := make([]*Colour, 0, 4)
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
		PrintColour(colour)
	}
}
