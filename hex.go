package main

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	hexDisabled = false
	hexColour   = regexp.MustCompile(fmt.Sprintf("(?:#|0x)(?:%s{8}|%[1]s{6}|%[1]s{3})", hexDigit))
)

func setTripleHexDisabled(disabled bool) {
	if disabled {
		hexColour = regexp.MustCompile(fmt.Sprintf("(?:#|0x)(?:%s{8}|%[1]s{6})", hexDigit))
	} else {
		hexColour = regexp.MustCompile(fmt.Sprintf("(?:#|0x)(?:%s{8}|%[1]s{6}|%[1]s{3})", hexDigit))
	}
}

func parseHex(line string) colours {
	var clrs colours
	if hexDisabled {
		return clrs
	}

	matches := hexColour.FindAllStringIndex(line, -1)
	for _, match := range matches {
		if !checkBoundary || isWord(line, match[0], match[1]) {
			colour := &Colour{
				ColStart: match[0] + 1,
				ColEnd:   match[1],
				Hex:      strings.ToLower(toFullHex(line[match[0]:match[1]])),
				Line:     line,
			}
			clrs = append(clrs, colour)
		}
	}
	return clrs
}
