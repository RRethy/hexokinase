package main

import (
	"bufio"
	"os"
	"sort"
)

// parseFile returns all colours matched when parsing in.
// tag is attached to each colour.
func parseFile(in *os.File, tag string, max int) colours {
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
				colour.Tag = tag
			}
			colours = append(colours, lineColours...)
		}
		if max != -1 && len(colours) >= max {
			break
		}
	}

	sort.Sort(colours)
	if max != -1 && len(colours) > max {
		colours = colours[:max]
	}

	return colours
}
