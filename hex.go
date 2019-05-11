package main

import (
	"fmt"
	"regexp"
)

const (
	hexDigit = "[0-9a-fA-F]"
)

var (
	hexPat = regexp.MustCompile(fmt.Sprintf("#(%s{6}|%[1]s{3})", hexDigit))
)

func parseLine(line string, lnum int) []colour {
	var colours []colour
	matches := hexPat.FindAllStringSubmatchIndex(line, -1)
	for _, match := range matches {
		colours = append(colours, colour{
			colStart: match[0],
			colEnd:   match[1],
			lnum:     lnum,
			hex:      toFullHex(line[match[2]:match[3]]),
		})
	}
	return colours
}

func toFullHex(str string) string {
	if len(str) == 6 {
		return str
	}
	return fmt.Sprintf("%c%c%c%c%c%c",
		str[0], str[0], str[1], str[1], str[2], str[2])
}
