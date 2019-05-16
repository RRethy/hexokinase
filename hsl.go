package main

import (
	"fmt"
	"regexp"
	"strconv"
)

const (
	validHue = "3(?:[0-5][0-9]|60)|[0-2]?[0-9]?[0-9]"
)

var (
	hslPat = regexp.MustCompile(fmt.Sprintf(`hsl\(\s*(%s)\s*,\s*(%s)\s*,\s*(%[2]s)\s*\)`, validHue, percentage))
)

func parseHSL(line string) []*Colour {
	var colours []*Colour
	matches := hslPat.FindAllStringSubmatchIndex(line, -1)
	for _, match := range matches {
		h, err := strconv.Atoi(line[match[2]:match[3]])
		s, err := strToDec(line[match[4]:match[5]])
		l, err := strToDec(line[match[4]:match[5]])
		if err != nil {
			continue
		}
		colour := &Colour{
			ColStart: match[0] + 1,
			ColEnd:   match[1],
			Hex:      hslToHex(h, s, l),
		}
		colours = append(colours, colour)
	}
	return colours
}
