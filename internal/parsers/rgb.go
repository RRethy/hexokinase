package parser

import (
	"fmt"
	"github.com/rrethy/hexokinase/internal/colour"
	"regexp"
	"strconv"
)

const (
	validNumber = "2(?:[0-4][0-9]|5[0-5])|1?[0-9]?[0-9]"
)

var (
	rgbPat = regexp.MustCompile(fmt.Sprintf(`rgb\(\s*(%s)\s*,\s*(%[1]s)\s*,\s*(%[1]s)\s*\)`, validNumber))
)

func parseRGB(line string, lnum int) []*colour.Colour {
	var colours []*colour.Colour
	matches := rgbPat.FindAllStringSubmatchIndex(line, -1)
	for _, match := range matches {
		colour := new(colour.Colour)
		colour.ColStart = match[0] + 1
		colour.ColEnd = match[1]
		colour.Lnum = lnum
		r, err := strconv.Atoi(line[match[2]:match[3]])
		g, err := strconv.Atoi(line[match[4]:match[5]])
		b, err := strconv.Atoi(line[match[6]:match[7]])
		if err != nil {
			continue
		}
		colour.Hex = rgbToHex(r, g, b)
		colours = append(colours, colour)
	}
	return colours
}

func rgbToHex(r, g, b int) string {
	return fmt.Sprintf("#%s%s%s",
		fmt.Sprintf("%02x", r),
		fmt.Sprintf("%02x", g),
		fmt.Sprintf("%02x", b))
}
