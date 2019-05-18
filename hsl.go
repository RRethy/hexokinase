package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	hslFunc     = regexp.MustCompile(fmt.Sprintf(`hsl\(\s*(%s)\s*,\s*(%s)\s*,\s*(%[2]s)\s*\)`, validHue, percentage))
	hslDisabled = false
)

func parseHSL(line string) colours {
	var clrs colours
	if hslDisabled {
		return clrs
	}

	matches := hslFunc.FindAllStringSubmatchIndex(line, -1)
	for _, match := range matches {
		h, err := strconv.ParseFloat(line[match[2]:match[3]], 64)
		s, err := percentageStrToInt(line[match[4]:match[5]])
		l, err := percentageStrToInt(line[match[6]:match[7]])
		if err != nil {
			continue
		}
		colour := &Colour{
			ColStart: match[0] + 1,
			ColEnd:   match[1],
			Hex:      hslToHex(float64(int(h)%360), float64(s)/100, float64(l)/100),
			Line:     line,
		}
		clrs = append(clrs, colour)
	}
	return clrs
}
