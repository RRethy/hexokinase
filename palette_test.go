package main

import (
	"regexp"
	"testing"
)

func TestPalette(t *testing.T) {
	p1 := &palette{
		ColourPairs: map[string]string{"foo": "bar", "BING": "BONG"},
	}
	p2 := &palette{
		ColourPairs:   map[string]string{"one": "1", "two": "2"},
		compiledRegex: regexp.MustCompile("one|two"),
	}
	palettes = append(palettes, p1, p2)
	var tests = []testData{
		{" nothing to see here", []*Colour{}},
		{"foo", []*Colour{
			&Colour{ColStart: 1, ColEnd: 3, Hex: "bar"},
		}},
		{"BING", []*Colour{
			&Colour{ColStart: 1, ColEnd: 4, Hex: "BONG"},
		}},
		{"one", []*Colour{
			&Colour{ColStart: 1, ColEnd: 3, Hex: "1"},
		}},
		{"two", []*Colour{
			&Colour{ColStart: 1, ColEnd: 3, Hex: "2"},
		}},
		{" two ", []*Colour{
			&Colour{ColStart: 2, ColEnd: 4, Hex: "2"},
		}},
		{" foo ", []*Colour{
			&Colour{ColStart: 2, ColEnd: 4, Hex: "bar"},
		}},
		{" foo one two BING ", []*Colour{
			&Colour{ColStart: 2, ColEnd: 4, Hex: "bar"},
			&Colour{ColStart: 14, ColEnd: 17, Hex: "BONG"},
			&Colour{ColStart: 6, ColEnd: 8, Hex: "1"},
			&Colour{ColStart: 10, ColEnd: 12, Hex: "2"},
		}},
	}
	runTests("TestPalette", t, tests, parsePalettes)
}
