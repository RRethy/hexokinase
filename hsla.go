package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	hslaFunc     = regexp.MustCompile(fmt.Sprintf(`hsla\(\s*(%s)\s*,\s*(%s)\s*,\s*(%[2]s)\s*,\s*(%s)\s*\)`, validHue, percentage, alphaPat))
	hslaDisabled = false
)

func parseHSLA(line string) colours {
	var clrs colours
	if hslaDisabled {
		return clrs
	}

	matches := hslaFunc.FindAllStringSubmatchIndex(line, -1)
	for _, match := range matches {
		h, err := strconv.ParseFloat(line[match[2]:match[3]], 64)
		s, err := percentageStrToInt(line[match[4]:match[5]])
		l, err := percentageStrToInt(line[match[6]:match[7]])
		alpha, err := strconv.ParseFloat(line[match[8]:match[9]], 64)
		if err != nil {
			continue
		}
		colour := &Colour{
			ColStart: match[0] + 1,
			ColEnd:   match[1],
			Hex:      hslaToHex(float64(int(h)%360), float64(s)/100, float64(l)/100, alpha),
			Line:     line,
		}
		clrs = append(clrs, colour)
	}
	return clrs
}
