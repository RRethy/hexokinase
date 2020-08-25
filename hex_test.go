package main

import (
	"testing"
)

func TestParseHex(t *testing.T) {
	var tests = []testData{
		// test various values
		{" #fff ", colours{
			&Colour{ColStart: 2, ColEnd: 5, Hex: "#ffffff"},
		}},
		{" #fff ", colours{
			&Colour{ColStart: 2, ColEnd: 5, Hex: "#ffffff"},
		}},
		{" 0xfff ", colours{
			&Colour{ColStart: 2, ColEnd: 6, Hex: "#ffffff"},
		}},
		{" #FFF ", colours{
			&Colour{ColStart: 2, ColEnd: 5, Hex: "#ffffff"},
		}},
		{" #ffffff ", colours{
			&Colour{ColStart: 2, ColEnd: 8, Hex: "#ffffff"},
		}},
		{" #FFFFFF ", colours{
			&Colour{ColStart: 2, ColEnd: 8, Hex: "#ffffff"},
		}},
		{"	#FFFFFF	", colours{
			&Colour{ColStart: 2, ColEnd: 8, Hex: "#ffffff"},
		}},
		{"	#FFFFFFFF ", colours{
			&Colour{ColStart: 2, ColEnd: 10, Hex: "#ffffff"},
		}},
		{"	#B0FD23E6 ", colours{
			&Colour{ColStart: 2, ColEnd: 10, Hex: "#b7fd38"},
		}},
		{" #A23f23 ", colours{
			&Colour{ColStart: 2, ColEnd: 8, Hex: "#a23f23"},
		}},
		{"#a8f9e9", colours{
			&Colour{ColStart: 1, ColEnd: 7, Hex: "#a8f9e9"},
		}},
		{"0xa8f9e9", colours{
			&Colour{ColStart: 1, ColEnd: 8, Hex: "#a8f9e9"},
		}},

		// test invalid values
		{" # fff  ", colours{}},
		{" #gggggg ", colours{}},
		{" #banana ", colours{}},
		{" banana ", colours{}},
		{" #ggg ", colours{}},

		// test multiple values
		{" #ae9 #A23f23 ", colours{
			&Colour{ColStart: 2, ColEnd: 5, Lnum: 0, Hex: "#aaee99"},
			&Colour{ColStart: 7, ColEnd: 13, Lnum: 0, Hex: "#a23f23"},
		}},
		{"#ae9 #A23f23", colours{
			&Colour{ColStart: 1, ColEnd: 4, Lnum: 0, Hex: "#aaee99"},
			&Colour{ColStart: 6, ColEnd: 12, Lnum: 0, Hex: "#a23f23"},
		}},
		{"#ae9 #A23f23 #000 #ae9#A23f23#000", colours{
			&Colour{ColStart: 1, ColEnd: 4, Lnum: 0, Hex: "#aaee99"},
			&Colour{ColStart: 6, ColEnd: 12, Lnum: 0, Hex: "#a23f23"},
			&Colour{ColStart: 14, ColEnd: 17, Lnum: 0, Hex: "#000000"},
			&Colour{ColStart: 19, ColEnd: 22, Lnum: 0, Hex: "#aaee99"},
			&Colour{ColStart: 23, ColEnd: 29, Lnum: 0, Hex: "#a23f23"},
			&Colour{ColStart: 30, ColEnd: 33, Lnum: 0, Hex: "#000000"},
		}},
	}
	runTests("TestParseHex", t, tests, parseHex)
}

func TestParseHexWithBoundary(t *testing.T) {
	checkBoundary = true
	var tests = []testData{
		// test various values
		{" #fff ", colours{
			&Colour{ColStart: 2, ColEnd: 5, Hex: "#ffffff"},
		}},
		{" #fff ", colours{
			&Colour{ColStart: 2, ColEnd: 5, Hex: "#ffffff"},
		}},
		{" #FFF ", colours{
			&Colour{ColStart: 2, ColEnd: 5, Hex: "#ffffff"},
		}},
		{" #ffffff ", colours{
			&Colour{ColStart: 2, ColEnd: 8, Hex: "#ffffff"},
		}},
		{" #FFFFFF ", colours{
			&Colour{ColStart: 2, ColEnd: 8, Hex: "#ffffff"},
		}},
		{"	#FFFFFF	", colours{
			&Colour{ColStart: 2, ColEnd: 8, Hex: "#ffffff"},
		}},
		{" #A23f23 ", colours{
			&Colour{ColStart: 2, ColEnd: 8, Hex: "#a23f23"},
		}},
		{"#a8f9e9", colours{
			&Colour{ColStart: 1, ColEnd: 7, Hex: "#a8f9e9"},
		}},

		// test invalid values
		{" # fff  ", colours{}},
		{" #gggggg ", colours{}},
		{" #banana ", colours{}},
		{" banana ", colours{}},
		{" #ggg ", colours{}},

		// test multiple values
		{" #ae9 #A23f23 ", colours{
			&Colour{ColStart: 2, ColEnd: 5, Lnum: 0, Hex: "#aaee99"},
			&Colour{ColStart: 7, ColEnd: 13, Lnum: 0, Hex: "#a23f23"},
		}},
		{"#ae9 #A23f23", colours{
			&Colour{ColStart: 1, ColEnd: 4, Lnum: 0, Hex: "#aaee99"},
			&Colour{ColStart: 6, ColEnd: 12, Lnum: 0, Hex: "#a23f23"},
		}},
		{"#ae9 #A23f23 #000 #ae9#A23f23#000", colours{
			&Colour{ColStart: 1, ColEnd: 4, Lnum: 0, Hex: "#aaee99"},
			&Colour{ColStart: 6, ColEnd: 12, Lnum: 0, Hex: "#a23f23"},
			&Colour{ColStart: 14, ColEnd: 17, Lnum: 0, Hex: "#000000"},
			&Colour{ColStart: 19, ColEnd: 22, Lnum: 0, Hex: "#aaee99"},
		}},
	}
	runTests("TestParseHex", t, tests, parseHex)
}
