package main

import (
	"testing"
)

func TestParseHSL(t *testing.T) {
	var tests = []testData{
		{" hsl(195, 100%, 50%) ", []*Colour{
			&Colour{ColStart: 2, ColEnd: 20, Hex: "#00bfff"},
		}},
		{" hsl(0, 0%, 100%) ", []*Colour{
			&Colour{ColStart: 2, ColEnd: 17, Hex: "#ffffff"},
		}},
		{" hsl(0, 0%, 0%) ", []*Colour{
			&Colour{ColStart: 2, ColEnd: 15, Hex: "#000000"},
		}},
		{" hsl( 0 , 0% , 0% ) ", []*Colour{
			&Colour{ColStart: 2, ColEnd: 19, Hex: "#000000"},
		}},
		{"hsl(0,0%,0%)", []*Colour{
			&Colour{ColStart: 1, ColEnd: 12, Hex: "#000000"},
		}},
		{"hsl(360,50%,50%)", []*Colour{
			&Colour{ColStart: 1, ColEnd: 16, Hex: "#bf3f3f"},
		}},
		{"hsl(500,50%,50%)", []*Colour{
			&Colour{ColStart: 1, ColEnd: 16, Hex: "#3fbf6a"},
		}},
		{"hsl(-500,50%,50%)", []*Colour{}},
		{"hsl(-500,500%,50%)", []*Colour{}},
		{"hsl(-500,50%,500%)", []*Colour{}},
	}

	runTests("TestParseHSL", t, tests, parseHSL)
}
