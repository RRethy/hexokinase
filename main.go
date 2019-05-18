package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	paletteUsage = `
palette file names.
This has the ability to define additional patterns to match and how to convert
them into hex values. The file must be a valid json file that looks like this:

{
  "regex_pattern": "foo[0-9]bar[0-9]baz[0-9]",
  "colour_table": {
    "foo1bar1baz1": "#eb00ff",
    "foo2bar2baz2": "#ffeb00",
    "foo3bar3baz3": "#00ffeb"
  }
}

The "regex_pattern" key is optional. If omitted, every key in the
"colour_table" map will be matched instead of using the regex.
Any key in the "colour_table" map will be matched and have the associated hex
string that is provided outputted.
If the regex matches a string which is not a key in the "colour_table" map, it
will be discarded as a false positive.
No checking is done on the hex strings so technically they can be any string.
`
)

var (
	paletteFnames    = flag.String("palettes", "", paletteUsage)
	fmtShort         = flag.Bool("simplified", false, "same as -extended but don't print the full line. Overrides -extended.")
	fmtExtended      = flag.Bool("extended", true, `print results in the format "filename:lnum:colstart-colend:hex:line"`)
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
