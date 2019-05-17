package main

import (
	"bufio"
	"os"
	"sort"
)

// parseFile returns all colours matched when parsing in
func parseFile(in *os.File) colours {
	scanner := bufio.NewScanner(in)
	colours := make(colours, 0, 4)
	parsers := []parser{
		parseHex,
		parseRGB,
		parseRGBA,
		parseHSL,
		parseHSLA,
		parsePalettes,
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

	sort.Sort(colours)

	return colours
}
