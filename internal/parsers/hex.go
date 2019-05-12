package parser

import (
	"fmt"
	"github.com/rrethy/hexokinase/internal/models"
	"regexp"
	"strings"
)

const (
	hexDigit = "[0-9a-fA-F]"
)

var (
	hexPat = regexp.MustCompile(fmt.Sprintf("#(?:%s{6}|%[1]s{3})", hexDigit))
)

func parseHex(line string, lnum int) []*models.Colour {
	var colours []*models.Colour
	matches := hexPat.FindAllStringIndex(line, -1)
	for _, match := range matches {
		colour := new(models.Colour)
		colour.ColStart = match[0] + 1
		colour.ColEnd = match[1]
		colour.Lnum = lnum
		colour.Hex = strings.ToLower(toFullHex(line[match[0]:match[1]]))
		colours = append(colours, colour)
	}
	return colours
}

func toFullHex(str string) string {
	if len(str) == 7 {
		return str
	}
	return fmt.Sprintf("#%c%c%c%c%c%c",
		str[1], str[1], str[2], str[2], str[3], str[3])
}
