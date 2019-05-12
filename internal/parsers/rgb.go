package parser

import (
	"fmt"
	"github.com/rrethy/hexokinase/internal/models"
	"os"
	"regexp"
	"strconv"
)

const (
	validNumber = "2(?:[0-4][0-9]|5[0-5])|1?[0-9]?[0-9]"
	percentage  = "1?[0-9]{0,2}%"
	value       = "(?:(?:" + validNumber + ")|(?:" + percentage + "))"
)

var (
	rgbPat = regexp.MustCompile(fmt.Sprintf(`rgb\(\s*(%s)\s*,\s*(%[1]s)\s*,\s*(%[1]s)\s*\)`, value))
)

func parseRGB(line string, lnum int) []*models.Colour {
	fmt.Fprintf(os.Stdout, "%s\n", value)
	var colours []*models.Colour
	matches := rgbPat.FindAllStringSubmatchIndex(line, -1)
	for _, match := range matches {
		colour := new(models.Colour)
		colour.ColStart = match[0] + 1
		colour.ColEnd = match[1]
		colour.Lnum = lnum
		r, err := strToDec(line[match[2]:match[3]])
		g, err := strToDec(line[match[4]:match[5]])
		b, err := strToDec(line[match[6]:match[7]])
		if err != nil {
			continue
		}
		colour.Hex = rgbToHex(r, g, b)
		colours = append(colours, colour)
	}
	return colours
}

func strToDec(str string) (int, error) {
	if str[len(str)-1] == '%' {
		num, err := strconv.Atoi(str[:len(str)-1])
		if err != nil {
			return 0, err
		}
		return num * 255 / 100, nil
	}
	return strconv.Atoi(str)
}

func rgbToHex(r, g, b int) string {
	return fmt.Sprintf("#%s%s%s",
		fmt.Sprintf("%02x", r),
		fmt.Sprintf("%02x", g),
		fmt.Sprintf("%02x", b))
}
