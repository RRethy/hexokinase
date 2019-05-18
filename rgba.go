package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var (
	rgbaFunc     = regexp.MustCompile(fmt.Sprintf(`rgba\(\s*(%s)\s*,\s*(%[1]s)\s*,\s*(%[1]s)\s*,\s*(%s)\s*\)`, rgbFuncParam, alphaPat))
	rgbaDisabled = false
)

func parseRGBA(line string) colours {
	var clrs colours
	if rgbaDisabled {
		return clrs
	}

	matches := rgbaFunc.FindAllStringSubmatchIndex(line, -1)
	for _, match := range matches {
		r, err := strToDec(line[match[2]:match[3]])
		g, err := strToDec(line[match[4]:match[5]])
		b, err := strToDec(line[match[6]:match[7]])
		alpha, err := strconv.ParseFloat(line[match[8]:match[9]], 64)
		if err != nil {
			continue
		}
		colour := &Colour{
			ColStart: match[0] + 1,
			ColEnd:   match[1],
			Hex:      rgbaToHex(r, g, b, alpha),
			Line:     line,
		}
		clrs = append(clrs, colour)
	}
	return clrs
}
