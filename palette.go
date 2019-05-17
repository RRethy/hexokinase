package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
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
)

func initializePalettes(fnames ...string) {
	for _, fname := range fnames {
		file, err := ioutil.ReadFile(fname)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s - %v\n", fname, err)
			continue
		}
		p := &palette{}
		err = json.Unmarshal(file, &p)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %v\n", fname, err)
			continue
		}
		p.compileRegex()
		palettes = append(palettes, p)
	}
}

func parsePalettes(line string) []*Colour {
	var colours []*Colour
	if len(palettes) == 0 {
		return colours
	}

	for _, p := range palettes {
		colours = append(colours, parsePalette(line, p)...)
	}

	return colours
}

func parsePalette(line string, p *palette) []*Colour {
	var colours []*Colour

	if p.compiledRegex != nil {
		matches := p.compiledRegex.FindAllStringIndex(line, -1)
		for _, match := range matches {
			name := line[match[0]:match[1]]
			if hex, ok := p.ColourPairs[name]; ok {
				colour := &Colour{
					ColStart: match[0] + 1,
					ColEnd:   match[1],
					Hex:      hex,
				}
				colours = append(colours, colour)
			}
		}
	} else {
		used := make([]bool, len(line))
		for name, hex := range p.ColourPairs {
			curLine := line
			for len(curLine) > 0 {
				offset := len(line) - len(curLine)
				index := strings.Index(curLine, name)
				if index != -1 {
					if !used[offset+index] {
						colour := &Colour{
							ColStart: offset + index + 1,
							ColEnd:   offset + index + len(name),
							Hex:      hex,
						}
						colours = append(colours, colour)
					}
					curLine = curLine[index+len(name):]
				} else {
					break
				}
			}
		}
	}

	return colours
}
