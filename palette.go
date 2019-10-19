package main

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"regexp"
	"strings"
)

type palette struct {
	Regex         string            `json:"regex_pattern"`
	ColourPairs   map[string]string `json:"colour_table"`
	compiledRegex *regexp.Regexp
}

func (p *palette) compileRegex() {
	if len(p.Regex) != 0 {
		p.compiledRegex = regexp.MustCompile(p.Regex)
	}
}

var (
	palettes []*palette

	// ErrIncompletePalette indicates that a palette file does not have the
	// required key "colour_table"
	ErrIncompletePalette = errors.New(`palette missing required "colour_table" key`)
)

// A PaletteError records an unsuccessful attempt to load a palette file
type PaletteError struct {
	Func  string // the failing function (LoadPalettes, etc.)
	Fname string //the file name of the palette
	Err   error  // the reason for failure
}

func (e *PaletteError) Error() string {
	return "hexokinase." + e.Func + ": " + "loading palette " + e.Fname + ": " + e.Err.Error()
}

// LoadPalettes reads the files in fnames and loads them as a palette if they are valid.
func LoadPalettes(fnames ...string) []error {
	const fnLoadPalettes = "LoadPalettes"

	var errs []error
	for _, fname := range fnames {
		file, err := ioutil.ReadFile(fname)
		if err != nil {
			errs = append(errs, &PaletteError{fnLoadPalettes, fname, err})
			continue
		}

		p := &palette{}
		err = json.Unmarshal(file, &p)
		if err != nil {
			errs = append(errs, &PaletteError{fnLoadPalettes, fname, err})
			continue
		}

		if len(p.ColourPairs) == 0 {
			errs = append(errs, &PaletteError{fnLoadPalettes, fname, ErrIncompletePalette})
			continue
		}

		p.compileRegex()
		palettes = append(palettes, p)
	}
	return errs
}

func parsePalettes(line string) colours {
	var clrs colours
	if len(palettes) == 0 {
		return clrs
	}

	for _, p := range palettes {
		clrs = append(clrs, parsePalette(line, p)...)
	}

	return clrs
}

func parsePalette(line string, p *palette) colours {
	var clrs colours

	if p.compiledRegex != nil {
		matches := p.compiledRegex.FindAllStringIndex(line, -1)
		for _, match := range matches {
			name := line[match[0]:match[1]]
			if hex, ok := p.ColourPairs[name]; ok {
				if !checkBoundary || isWord(line, match[0], match[1]) {
					colour := &Colour{
						ColStart: match[0] + 1,
						ColEnd:   match[1],
						Hex:      hex,
						Line:     line,
					}
					clrs = append(clrs, colour)
				}
			}
		}
	} else {
		for name, hex := range p.ColourPairs {
			curLine := line
			for len(curLine) > 0 {
				offset := len(line) - len(curLine)
				index := strings.Index(curLine, name)
				if index != -1 {
					if !checkBoundary || isWord(line, offset+index, offset+index+len(name)) {
						colour := &Colour{
							ColStart: offset + index + 1,
							ColEnd:   offset + index + len(name),
							Hex:      hex,
							Line:     line,
						}
						clrs = append(clrs, colour)
					}
					curLine = curLine[index+len(name):]
				} else {
					break
				}
			}
		}
	}

	return clrs
}
