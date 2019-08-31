package main

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	hexDisabled   = false
	hexColour     = regexp.MustCompile(fmt.Sprintf("#(?:%s{6}|%[1]s{3})", hexDigit))
	checkBoundary = false
)

func setTripleHexDisabled(disabled bool) {
	if disabled {
		hexColour = regexp.MustCompile(fmt.Sprintf("#%s{6}", hexDigit))
	} else {
		hexColour = regexp.MustCompile(fmt.Sprintf("#(?:%s{6}|%[1]s{3})", hexDigit))
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

// poor mans vimscript "\<\>"
func isWord(line string, start int, end int) bool {
	return (start == 0 || !isKeyword(line[start-1])) && (end == len(line) || !isKeyword(line[end]))
}

// TODO do a better job with utf-8
func isKeyword(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '_' || c == '-'
}
