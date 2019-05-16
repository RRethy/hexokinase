package main

import (
	"bufio"
	"os"
	"testing"
)

func TestParseWebColours(t *testing.T) {
	var tests = []testData{
		// test various values
		{"aliceblue", []*Colour{
			&Colour{ColStart: 1, ColEnd: 9, Hex: "#f0f8ff"},
		}},
		{"mediumseagreen", []*Colour{
			&Colour{ColStart: 1, ColEnd: 14, Hex: "#3cb371"},
		}},
		{"darkolivegreen", []*Colour{
			&Colour{ColStart: 1, ColEnd: 14, Hex: "#556b2f"},
		}},
		{"cyan", []*Colour{
			&Colour{ColStart: 1, ColEnd: 4, Hex: "#00ffff"},
		}},
		{"blue", []*Colour{
			&Colour{ColStart: 1, ColEnd: 4, Hex: "#0000ff"},
		}},
		{"grey", []*Colour{
			&Colour{ColStart: 1, ColEnd: 4, Hex: "#808080"},
		}},
		{"gray", []*Colour{
			&Colour{ColStart: 1, ColEnd: 4, Hex: "#808080"},
		}},
		{" gre y ", []*Colour{}},
		{" seven ", []*Colour{}},
		// TODO unicode
		// {"  …grey… ", []*Colour{
		// 	&Colour{ColStart: 1, ColEnd: 4, Hex: "#808080"},
		// }},
		{"grey graygrey", []*Colour{
			&Colour{ColStart: 1, ColEnd: 4, Hex: "#808080"},
			&Colour{ColStart: 10, ColEnd: 13, Hex: "#808080"},
			&Colour{ColStart: 6, ColEnd: 9, Hex: "#808080"},
		}},
		{"grey aliceblue blue", []*Colour{
			&Colour{ColStart: 6, ColEnd: 14, Hex: "#f0f8ff"},
			&Colour{ColStart: 16, ColEnd: 19, Hex: "#0000ff"},
			&Colour{ColStart: 1, ColEnd: 4, Hex: "#808080"},
		}},
		{" 	  	grey aliceblue   blue", []*Colour{
			&Colour{ColStart: 11, ColEnd: 19, Hex: "#f0f8ff"},
			&Colour{ColStart: 23, ColEnd: 26, Hex: "#0000ff"},
			&Colour{ColStart: 6, ColEnd: 9, Hex: "#808080"},
		}},
	}
	runTests("TestParseWebColours", t, tests, parseWebColours)
}

func BenchmarkParseWebColours(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, err := os.Open("./test_colours.txt")
		if err != nil {
			b.Errorf("%v\n", err)
			continue
		}
		scanner := bufio.NewScanner(file)
		if scanner.Scan() {
			parseWebColours(scanner.Text())
		}
		file.Close()
	}
}
