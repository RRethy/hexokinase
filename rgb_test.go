package main

import (
	"testing"
)

func TestParseRGB(t *testing.T) {
	var tests = []testData{
		{"rgb(0,0,0)", colours{
			&Colour{ColStart: 1, ColEnd: 10, Hex: "#000000"},
		}},
		{"…rgb(0,0,0)", colours{
			&Colour{ColStart: 4, ColEnd: 13, Hex: "#000000"},
		}},
		{"rgb(176,253,35)", colours{
			&Colour{ColStart: 1, ColEnd: 15, Hex: "#b0fd23"},
		}},
		{"rgb(176.99,253.12,35.99)", colours{
			&Colour{ColStart: 1, ColEnd: 24, Hex: "#b0fd23"},
		}},

		// test percentages
		{"rgb(0%,253,35)", colours{
			&Colour{ColStart: 1, ColEnd: 14, Hex: "#00fd23"},
		}},
		{"rgb(100%,253,35)", colours{
			&Colour{ColStart: 1, ColEnd: 16, Hex: "#fffd23"},
		}},
		{"rgb(25%,253,35)", colours{
			&Colour{ColStart: 1, ColEnd: 15, Hex: "#3ffd23"},
		}},
		{"rgb(253,25%,35)", colours{
			&Colour{ColStart: 1, ColEnd: 15, Hex: "#fd3f23"},
		}},
		{"rgb(35,253,25%)", colours{
			&Colour{ColStart: 1, ColEnd: 15, Hex: "#23fd3f"},
		}},
		{"rgb(0%,25%,35%)", colours{
			&Colour{ColStart: 1, ColEnd: 15, Hex: "#003f59"},
		}},

		// test red value
		{"rgb(35,0,0)", colours{
			&Colour{ColStart: 1, ColEnd: 11, Hex: "#230000"},
		}},
		{"rgb(176,0,0)", colours{
			&Colour{ColStart: 1, ColEnd: 12, Hex: "#b00000"},
		}},
		{"rgb(215,0,0)", colours{
			&Colour{ColStart: 1, ColEnd: 12, Hex: "#d70000"},
		}},
		{"rgb(253,0,0)", colours{
			&Colour{ColStart: 1, ColEnd: 12, Hex: "#fd0000"},
		}},
		{"rgb(255,0,0)", colours{
			&Colour{ColStart: 1, ColEnd: 12, Hex: "#ff0000"},
		}},

		// test green value
		{"rgb(0,35,0)", colours{
			&Colour{ColStart: 1, ColEnd: 11, Hex: "#002300"},
		}},
		{"rgb(0,176,0)", colours{
			&Colour{ColStart: 1, ColEnd: 12, Hex: "#00b000"},
		}},
		{"rgb(0,215,0)", colours{
			&Colour{ColStart: 1, ColEnd: 12, Hex: "#00d700"},
		}},
		{"rgb(0,253,0)", colours{
			&Colour{ColStart: 1, ColEnd: 12, Hex: "#00fd00"},
		}},
		{"rgb(0,255,0)", colours{
			&Colour{ColStart: 1, ColEnd: 12, Hex: "#00ff00"},
		}},

		// test blue value
		{"rgb(0,0,35)", colours{
			&Colour{ColStart: 1, ColEnd: 11, Hex: "#000023"},
		}},
		{"rgb(0,0,176)", colours{
			&Colour{ColStart: 1, ColEnd: 12, Hex: "#0000b0"},
		}},
		{"rgb(0,0,215)", colours{
			&Colour{ColStart: 1, ColEnd: 12, Hex: "#0000d7"},
		}},
		{"rgb(0,0,253)", colours{
			&Colour{ColStart: 1, ColEnd: 12, Hex: "#0000fd"},
		}},
		{"rgb(0,0,255)", colours{
			&Colour{ColStart: 1, ColEnd: 12, Hex: "#0000ff"},
		}},

		// test multiple values
		{"rgb(0,0,255)rgb(176,253,35)  rgb(176,253,35)", colours{
			&Colour{ColStart: 1, ColEnd: 12, Hex: "#0000ff"},
			&Colour{ColStart: 13, ColEnd: 27, Hex: "#b0fd23"},
			&Colour{ColStart: 30, ColEnd: 44, Hex: "#b0fd23"},
		}},

		// tests invalid values
		{"rgb(256,0,0)", colours{}},
		{"rgb(0,0,256)", colours{}},
		{"rgb(0,0,256)", colours{}},
		{"rgb(1000,1000,1000)", colours{}},

		// test handling of whitespace
		{" rgb(0,0,0) ", colours{
			&Colour{ColStart: 2, ColEnd: 11, Hex: "#000000"},
		}},
		{" rgb(0,0,0) rgb(0,0,0) ", colours{
			&Colour{ColStart: 2, ColEnd: 11, Hex: "#000000"},
			&Colour{ColStart: 13, ColEnd: 22, Hex: "#000000"},
		}},
		{"rgb( 0 , 0 , 0 )", colours{
			&Colour{ColStart: 1, ColEnd: 16, Hex: "#000000"},
		}},
		{"rgb(  0  ,  0  ,  0  )", colours{
			&Colour{ColStart: 1, ColEnd: 22, Hex: "#000000"},
		}},
		{"rgb(	0	,	0	,	0	)", colours{
			&Colour{ColStart: 1, ColEnd: 16, Hex: "#000000"},
		}},
	}
	runTests("TestParseRGB", t, tests, parseRGB)
}
