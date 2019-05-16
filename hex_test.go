package main

import (
	"testing"
)

type colours []*Colour

func TestParseHex(t *testing.T) {
	var tests = []testData{
		// test various values
		{" #fff ", []*Colour{
			&Colour{ColStart: 2, ColEnd: 5, Hex: "#ffffff"},
		}},
		{" #fff ", []*Colour{
			&Colour{ColStart: 2, ColEnd: 5, Hex: "#ffffff"},
		}},
		{" #FFF ", []*Colour{
			&Colour{ColStart: 2, ColEnd: 5, Hex: "#ffffff"},
		}},
		{" #ffffff ", []*Colour{
			&Colour{ColStart: 2, ColEnd: 8, Hex: "#ffffff"},
		}},
		{" #FFFFFF ", []*Colour{
			&Colour{ColStart: 2, ColEnd: 8, Hex: "#ffffff"},
		}},
		{"	#FFFFFF	", []*Colour{
			&Colour{ColStart: 2, ColEnd: 8, Hex: "#ffffff"},
		}},
		{" #A23f23 ", []*Colour{
			&Colour{ColStart: 2, ColEnd: 8, Hex: "#a23f23"},
		}},
		{"#a8f9e9", []*Colour{
			&Colour{ColStart: 1, ColEnd: 7, Hex: "#a8f9e9"},
		}},

		// test invalid values
		{" # fff  ", []*Colour{}},
		{" #gggggg ", []*Colour{}},
		{" #banana ", []*Colour{}},
		{" banana ", []*Colour{}},
		{" #ggg ", []*Colour{}},

		// test multiple values
		{" #ae9 #A23f23 ", []*Colour{
			&Colour{ColStart: 2, ColEnd: 5, Lnum: 0, Hex: "#aaee99"},
			&Colour{ColStart: 7, ColEnd: 13, Lnum: 0, Hex: "#a23f23"},
		}},
		{"#ae9#A23f23", []*Colour{
			&Colour{ColStart: 1, ColEnd: 4, Lnum: 0, Hex: "#aaee99"},
			&Colour{ColStart: 5, ColEnd: 11, Lnum: 0, Hex: "#a23f23"},
		}},
		{"#ae9#A23f23#000 #ae9#A23f23#000", []*Colour{
			&Colour{ColStart: 1, ColEnd: 4, Lnum: 0, Hex: "#aaee99"},
			&Colour{ColStart: 5, ColEnd: 11, Lnum: 0, Hex: "#a23f23"},
			&Colour{ColStart: 12, ColEnd: 15, Lnum: 0, Hex: "#000000"},
			&Colour{ColStart: 17, ColEnd: 20, Lnum: 0, Hex: "#aaee99"},
			&Colour{ColStart: 21, ColEnd: 27, Lnum: 0, Hex: "#a23f23"},
			&Colour{ColStart: 28, ColEnd: 31, Lnum: 0, Hex: "#000000"},
		}},
	}
	runTests("TestParseHex", t, tests, parseHex)
}
