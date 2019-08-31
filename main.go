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
	disabledPatterns = flag.String("dp", "", "disabled patterns which will not be parsed for. Comma separated list\nwith possible values of full_hex, triple_hex, rgb, rgba, hsl, hsla, colour_names. The \"names\"\nargument refers to web colour names.")
	enabledPatterns  = flag.String("ep", "", "enabled patterns which will be parsed for. Comma separated list\nwith possible values of full_hex, triple_hex, rgb, rgba, hsl, hsla, colour_names. The \"names\"\nargument refers to web colour names.")
	fnames           = flag.String("files", "stdin", "files to parse (or stdin to parse stdin)")
	reverse          = flag.Bool("r", false, "reverse output")
	checkForColour   = flag.String("check", "", "file to check if it contains colour patterns. This will override -fnames. A non-zero exit status indicates no colours found.")
	bgHex            = flag.String("bg", "#ffffff", "background colour used for alpha calculations with rgba and hsla functions.")
	useBoundaries    = flag.Bool("boundary", false, "TODO")
)

func main() {
	flag.Parse()

	SetBgHex(*bgHex)
	checkBoundary = *useBoundaries

	if len(*paletteFnames) > 0 {
		errs := LoadPalettes(strings.Split(*paletteFnames, ",")...)
		if len(errs) > 0 {
			for _, err := range errs {
				fmt.Fprintf(os.Stderr, "%v\n", err)
			}
		}
	}

	if len(*disabledPatterns) > 0 {
		for _, pattern := range strings.Split(*disabledPatterns, ",") {
			if len(pattern) == 0 {
				continue
			}

			switch pattern {
			case "full_hex":
				hexDisabled = true
			case "rgb":
				rgbDisabled = true
			case "rgba":
				rgbaDisabled = true
			case "hsl":
				hslDisabled = true
			case "hsla":
				hslaDisabled = true
			case "colour_names":
				webColoursDisabled = true
			case "triple_hex":
				setTripleHexDisabled(true)
			default:
				fmt.Fprintf(os.Stderr, "Unknown argument to flag -dp: %s", pattern)
			}
		}
	} else if len(*enabledPatterns) > 0 {
		hexDisabled = true
		rgbDisabled = true
		rgbaDisabled = true
		hslDisabled = true
		hslaDisabled = true
		webColoursDisabled = true
		setTripleHexDisabled(true)
		for _, pattern := range strings.Split(*enabledPatterns, ",") {
			if len(pattern) == 0 {
				continue
			}

			switch pattern {
			case "full_hex":
				hexDisabled = false
			case "rgb":
				rgbDisabled = false
			case "rgba":
				rgbaDisabled = false
			case "hsl":
				hslDisabled = false
			case "hsla":
				hslaDisabled = false
			case "colour_names":
				webColoursDisabled = false
			case "triple_hex":
				setTripleHexDisabled(false)
			default:
				fmt.Fprintf(os.Stderr, "Unknown argument to flag -ep: %s", pattern)
			}
		}
	}

	if *fmtShort {
		SetOutputFmt(ShortFmt)
	} else {
		SetOutputFmt(ExtendedFmt)
	}

	if len(*checkForColour) > 0 {
		file, err := os.Open(*checkForColour)
		defer file.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		clrs := parseFile(file, *checkForColour, 1)
		if len(clrs) > 0 {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	} else {
		for _, fname := range strings.Split(*fnames, ",") {
			var file *os.File
			var err error

			if fname == "stdin" {
				file = os.Stdin
			} else {
				file, err = os.Open(fname)
				defer file.Close()
			}
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}

			clrs := parseFile(file, fname, -1)
			PrintColours(clrs, os.Stdout, *reverse)
		}
	}
}
