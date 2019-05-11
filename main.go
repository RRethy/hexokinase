package main

import (
	"bufio"
	"fmt"
	"os"
)

type colour struct {
	colStart, colEnd int
	lnum             int
	hex              string
}

func main() {
	if err := readInput(); err != nil {
		fmt.Fprintf(os.Stderr, "reading Stdin failed: %v", err)
	}
}

func readInput() error {
	scanner := bufio.NewScanner(os.Stdin)
	lineNum := 1
	for scanner.Scan() {
		colours := parseLine(scanner.Text(), lineNum)
		if len(colours) != 0 {
			for _, colour := range colours {
				fmt.Fprintf(os.Stdout, "%d:%d-%d:%s\n", colour.lnum, colour.colStart, colour.colEnd, colour.hex)
			}
		}
		lineNum++
	}
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

// func printMatch(match colour) {
// }

// func parseLine(line string, lineNum int) {
// 6-3 digit hex
// hexDigit := "[0-9a-fA-F]"
// hexPat := regexp.MustCompile(fmt.Sprintf("%s{6}|%s{3}", hexDigit, hexDigit))
// matches := hexPat.FindAllStringIndex(line, -1)
// for _, tuple := range matches {
// 	fmt.Fprintf(os.Stdout, "%d:%d-%d\n", lineNum, tuple[0], tuple[1])
// }

// rgb
// validNumber := "(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)"
// rgbPat := regexp.MustCompile(fmt.Sprintf(`rgb\(\s*%s\s*,\s*%s\s*,\s*%s\s*\)`,
// 	validNumber, validNumber, validNumber))
// matches := rgbPat.FindAllStringSubmatchIndex(line, -1)
// for _, match := range matches {
// 	r, err := strconv.Atoi(line[match[2]:match[3]])
// 	g, err := strconv.Atoi(line[match[4]:match[5]])
// 	b, err := strconv.Atoi(line[match[6]:match[7]])
// 	if err != nil {
// 		continue
// 	}
// 	fmt.Fprintf(os.Stdout, "%d:%d-%d:(%d, %d, %d)\n",
// 		lineNum, match[0], match[1], r, g, b)
// }
// }
