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

func parseHex(line string) []*models.Colour {
	var colours []*models.Colour
	matches := hexPat.FindAllStringIndex(line, -1)
	for _, match := range matches {
		colour := &models.Colour{
			ColStart: match[0] + 1,
			ColEnd:   match[1],
			Hex:      strings.ToLower(toFullHex(line[match[0]:match[1]])),
		}
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
