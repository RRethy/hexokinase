package main

import (
	"testing"
)

func TestParseHSLA(t *testing.T) {
	var tests = []testData{
		{" hsla(195, 100%, 50%, 0) ", colours{
			&Colour{ColStart: 2, ColEnd: 24, Hex: "#ffffff"},
		}},
		{" hsla(195.5, 100%, 50%, 0) ", colours{
			&Colour{ColStart: 2, ColEnd: 26, Hex: "#ffffff"},
		}},
		{" hsla(195.53, 100%, 50%, 0) ", colours{
			&Colour{ColStart: 2, ColEnd: 27, Hex: "#ffffff"},
		}},
		{" hsla(195, 100%, 50%, 0.5) ", colours{
			&Colour{ColStart: 2, ColEnd: 26, Hex: "#7fdfff"},
		}},
		{" hsla(195, 100%, 50%, 1) ", colours{
			&Colour{ColStart: 2, ColEnd: 24, Hex: "#00bfff"},
		}},
		{" hsla(0, 0%, 100%, 1) ", colours{
			&Colour{ColStart: 2, ColEnd: 21, Hex: "#ffffff"},
		}},
		{" hsla(0, 0%, 0%, 1) ", colours{
			&Colour{ColStart: 2, ColEnd: 19, Hex: "#000000"},
		}},
		{" hsla( 0 , 0% , 0% , 1 ) ", colours{
			&Colour{ColStart: 2, ColEnd: 24, Hex: "#000000"},
		}},
		{"hsla(0,0%,0%, 1)", colours{
			&Colour{ColStart: 1, ColEnd: 16, Hex: "#000000"},
		}},
		{"hsla(360,50%,50%, 1)", colours{
			&Colour{ColStart: 1, ColEnd: 20, Hex: "#bf3f3f"},
		}},
		{"hsla(500,50%,50%, 1)", colours{
			&Colour{ColStart: 1, ColEnd: 20, Hex: "#3fbf6a"},
		}},
		{"hsla(-500,50%,50%, 1)", colours{}},
		{"hsla(-500,500%,50%, 1)", colours{}},
		{"hsla(-500,50%,500%, 1)", colours{}},
		{"hsla(195.531, 100%, 50%, 0)", colours{}},
	}

	runTests("TestParseHSLA", t, tests, parseHSLA)
}
