package main

import (
	"regexp"
	"testing"
)

func TestPalette(t *testing.T) {
	p1 := &palette{
		colourPairs: map[string]string{"foo": "bar", "BING": "BONG"},
	}
	p2 := &palette{
		colourPairs:   map[string]string{"one": "1", "two": "2"},
		compiledRegex: regexp.MustCompile("one|two"),
	}
	p3 := &palette{
		colourPairs:   map[string]string{},
		compiledRegex: regexp.MustCompile("three"),
	}
	palettes = append(palettes, p1, p2, p3)
	var tests = []testData{
		{" nothing to see here", []Colour{}},
		{" three ", []Colour{}},
		{"foo", []Colour{
			Colour{ColStart: 1, ColEnd: 3, Hex: "bar"},
		}},
		{"BING", []Colour{
			Colour{ColStart: 1, ColEnd: 4, Hex: "BONG"},
		}},
		{"one", []Colour{
			Colour{ColStart: 1, ColEnd: 3, Hex: "1"},
		}},
		{"two", []Colour{
			Colour{ColStart: 1, ColEnd: 3, Hex: "2"},
		}},
		{" two ", []Colour{
			Colour{ColStart: 2, ColEnd: 4, Hex: "2"},
		}},
		{" foo ", []Colour{
			Colour{ColStart: 2, ColEnd: 4, Hex: "bar"},
		}},
		{" foo two ", []Colour{
			Colour{ColStart: 2, ColEnd: 4, Hex: "bar"},
			Colour{ColStart: 6, ColEnd: 8, Hex: "2"},
		}},
	}
	runTests("TestPalette", t, tests, parsePalettes)
}
