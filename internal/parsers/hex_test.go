package parser

import (
	"github.com/rrethy/hexokinase/internal/colour"
	"testing"
)

type colours []*colour.Colour

func TestParseHex(t *testing.T) {
	var tests = []struct {
		line string
		lnum int
		want colours
	}{
		// test various values
		{" #fff ", 0, []*colour.Colour{
			&colour.Colour{ColStart: 2, ColEnd: 5, Lnum: 0, Hex: "#ffffff"},
		}},
		{" #fff ", 1, []*colour.Colour{
			&colour.Colour{ColStart: 2, ColEnd: 5, Lnum: 1, Hex: "#ffffff"},
		}},
		{" #FFF ", 0, []*colour.Colour{
			&colour.Colour{ColStart: 2, ColEnd: 5, Lnum: 0, Hex: "#ffffff"},
		}},
		{" #ffffff ", 0, []*colour.Colour{
			&colour.Colour{ColStart: 2, ColEnd: 8, Lnum: 0, Hex: "#ffffff"},
		}},
		{" #FFFFFF ", 0, []*colour.Colour{
			&colour.Colour{ColStart: 2, ColEnd: 8, Lnum: 0, Hex: "#ffffff"},
		}},
		{"	#FFFFFF	", 0, []*colour.Colour{
			&colour.Colour{ColStart: 2, ColEnd: 8, Lnum: 0, Hex: "#ffffff"},
		}},
		{" #A23f23 ", 0, []*colour.Colour{
			&colour.Colour{ColStart: 2, ColEnd: 8, Lnum: 0, Hex: "#a23f23"},
		}},
		{"#a8f9e9", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 7, Lnum: 0, Hex: "#a8f9e9"},
		}},

		// test invalid values
		{" # fff  ", 0, []*colour.Colour{}},
		{" #gggggg ", 0, []*colour.Colour{}},
		{" #banana ", 0, []*colour.Colour{}},
		{" banana ", 0, []*colour.Colour{}},
		{" #ggg ", 0, []*colour.Colour{}},

		// test multiple values
		{" #ae9 #A23f23 ", 0, []*colour.Colour{
			&colour.Colour{ColStart: 2, ColEnd: 5, Lnum: 0, Hex: "#aaee99"},
			&colour.Colour{ColStart: 7, ColEnd: 13, Lnum: 0, Hex: "#a23f23"},
		}},
		{"#ae9#A23f23", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 4, Lnum: 0, Hex: "#aaee99"},
			&colour.Colour{ColStart: 5, ColEnd: 11, Lnum: 0, Hex: "#a23f23"},
		}},
		{"#ae9#A23f23#000 #ae9#A23f23#000", 0, []*colour.Colour{
			&colour.Colour{ColStart: 1, ColEnd: 4, Lnum: 0, Hex: "#aaee99"},
			&colour.Colour{ColStart: 5, ColEnd: 11, Lnum: 0, Hex: "#a23f23"},
			&colour.Colour{ColStart: 12, ColEnd: 15, Lnum: 0, Hex: "#000000"},
			&colour.Colour{ColStart: 17, ColEnd: 20, Lnum: 0, Hex: "#aaee99"},
			&colour.Colour{ColStart: 21, ColEnd: 27, Lnum: 0, Hex: "#a23f23"},
			&colour.Colour{ColStart: 28, ColEnd: 31, Lnum: 0, Hex: "#000000"},
		}},
	}
	for i, test := range tests {
		if got := parseHex(test.line, test.lnum); !areSameColours(got, test.want) {
			t.Errorf("test number: %d\n\tgot:    %v\n\twanted: %v", i, got, test.want)
		}
	}
}
