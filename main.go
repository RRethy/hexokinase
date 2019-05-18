package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	paletteFnames    = flag.String("palettes", "", "palette file names")
	fmtShort         = flag.Bool("s", false, "same as -e but don't print the full line. Overrides -e.")
	fmtExtended      = flag.Bool("e", true, `print results in the format "filename:lnum:colstart-colend:hex:line"`)
	disabledPatterns = flag.String("dp", "", "disabled patterns which will not be parsed for. Comma separated list\nwith possible values of hex, rgb, rgba, hsl, hsla, names. The \"names\"\nargument refers to web colour names.")
	fnames           = flag.String("files", "stdout", "files to parse (or stdout to parse stdout)")
)

func main() {
	flag.Parse()

	if len(*paletteFnames) > 0 {
		errs := LoadPalettes(strings.Split(*paletteFnames, ",")...)
		if len(errs) > 0 {
			for _, err := range errs {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
		}
	}

	for _, pattern := range strings.Split(*disabledPatterns, ",") {
		switch pattern {
		case "hex":
			hexDisabled = true
		case "rgb":
			rgbDisabled = true
		case "rgba":
			rgbaDisabled = true
		case "hsl":
			hslDisabled = true
		case "hsla":
			hslaDisabled = true
		case "names":
			webColoursDisabled = true
		default:
			fmt.Fprintf(os.Stderr, "Unknown argument to flag -dp: %s", pattern)
		}
	}

	if *fmtShort {
		SetOutputFmt(ShortFmt)
	} else {
		SetOutputFmt(ExtendedFmt)
	}

	for _, fname := range strings.Split(*fnames, ",") {
		var file *os.File
		var err error

		if fname == "stdout" {
			file = os.Stdout
		} else {
			file, err = os.Open(fname)
			defer file.Close()
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			continue
		}

		clrs := parseFile(file, fname)
		PrintColours(clrs, os.Stdout)
	}
}
