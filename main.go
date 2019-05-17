package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	paletteFnames = flag.String("palettes", "", "palette file names")
)

func main() {
	flag.Parse()

	errs := LoadPalettes(strings.Split(*paletteFnames, ",")...)
	if errs != nil {
		for _, err := range errs {
			fmt.Fprintf(os.Stderr, "%v\n", err)
		}
	}

	SetOut(os.Stdout)
	clrs := parseFile(os.Stdin)
	printColours(clrs)
}
