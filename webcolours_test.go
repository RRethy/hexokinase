package main

import (
	"bufio"
	"os"
	"testing"
)

func TestParseWebColours(t *testing.T) {
	checkBoundary = false
	var tests = []testData{
		// test various values
		{"aliceblue", colours{
			&Colour{ColStart: 1, ColEnd: 9, Hex: "#f0f8ff"},
		}},
		{"mediumseagreen", colours{
			&Colour{ColStart: 1, ColEnd: 14, Hex: "#3cb371"},
		}},
		{"darkolivegreen", colours{
			&Colour{ColStart: 1, ColEnd: 14, Hex: "#556b2f"},
		}},
		{"cyan", colours{
			&Colour{ColStart: 1, ColEnd: 4, Hex: "#00ffff"},
		}},
		{"blue", colours{
			&Colour{ColStart: 1, ColEnd: 4, Hex: "#0000ff"},
		}},
		{"grey", colours{
			&Colour{ColStart: 1, ColEnd: 4, Hex: "#808080"},
		}},
		{"gray", colours{
			&Colour{ColStart: 1, ColEnd: 4, Hex: "#808080"},
		}},
		{" gre y ", colours{}},
		{" seven ", colours{}},
		{"  …grey… ", colours{
			&Colour{ColStart: 6, ColEnd: 9, Hex: "#808080"},
		}},
		{"grey graygrey", colours{
			&Colour{ColStart: 1, ColEnd: 4, Hex: "#808080"},
			&Colour{ColStart: 10, ColEnd: 13, Hex: "#808080"},
			&Colour{ColStart: 6, ColEnd: 9, Hex: "#808080"},
		}},
		{"grey aliceblue blue", colours{
			&Colour{ColStart: 6, ColEnd: 14, Hex: "#f0f8ff"},
			&Colour{ColStart: 16, ColEnd: 19, Hex: "#0000ff"},
			&Colour{ColStart: 1, ColEnd: 4, Hex: "#808080"},
		}},
		{" 	  	grey aliceblue   blue", colours{
			&Colour{ColStart: 11, ColEnd: 19, Hex: "#f0f8ff"},
			&Colour{ColStart: 23, ColEnd: 26, Hex: "#0000ff"},
			&Colour{ColStart: 6, ColEnd: 9, Hex: "#808080"},
		}},
	}
	runTests("TestParseWebColours", t, tests, parseWebColours)
}

func BenchmarkParseWebColours(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, err := os.Open("./benchmark_web_colours.txt")
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
