package main

import (
	"fmt"
	"regexp"
)

var (
	rgbFunc     = regexp.MustCompile(fmt.Sprintf(`rgb\(\s*(%s)\s*,\s*(%[1]s)\s*,\s*(%[1]s)\s*\)`, rgbFuncParam))
	rgbDisabled = false
)

func parseRGB(line string) colours {
	var clrs colours
	if rgbDisabled {
		return clrs
	}

	matches := rgbFunc.FindAllStringSubmatchIndex(line, -1)
	for _, match := range matches {
		r, err := strToDec(line[match[2]:match[3]])
		g, err := strToDec(line[match[4]:match[5]])
		b, err := strToDec(line[match[6]:match[7]])
		if err != nil {
			continue
		}
		colour := &Colour{
			ColStart: match[0] + 1,
			ColEnd:   match[1],
			Hex:      rgbToHex(r, g, b),
			Line:     line,
		}
		clrs = append(clrs, colour)
	}
	return clrs
}
