package parser

import (
	"bufio"
	"github.com/rrethy/hexokinase/internal/models"
	"os"
	"testing"
)

func TestParseWebColours(t *testing.T) {
	var tests = []testData{
		// test various values
		{"aliceblue", []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 9, Hex: "#f0f8ff"},
		}},
		{"mediumseagreen", []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 14, Hex: "#3cb371"},
		}},
		{"darkolivegreen", []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 14, Hex: "#556b2f"},
		}},
		{"cyan", []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 4, Hex: "#00ffff"},
		}},
		{"blue", []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 4, Hex: "#0000ff"},
		}},
		{"grey", []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 4, Hex: "#808080"},
		}},
		{"gray", []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 4, Hex: "#808080"},
		}},
		{" gre y ", []*models.Colour{}},
		{" seven ", []*models.Colour{}},
		// TODO unicode
		// {"  …grey… ", []*models.Colour{
		// 	&models.Colour{ColStart: 1, ColEnd: 4, Hex: "#808080"},
		// }},
		{"grey graygrey", []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 4, Hex: "#808080"},
			&models.Colour{ColStart: 10, ColEnd: 13, Hex: "#808080"},
			&models.Colour{ColStart: 6, ColEnd: 9, Hex: "#808080"},
		}},
		{"grey aliceblue blue", []*models.Colour{
			&models.Colour{ColStart: 6, ColEnd: 14, Hex: "#f0f8ff"},
			&models.Colour{ColStart: 16, ColEnd: 19, Hex: "#0000ff"},
			&models.Colour{ColStart: 1, ColEnd: 4, Hex: "#808080"},
		}},
		{" 	  	grey aliceblue   blue", []*models.Colour{
			&models.Colour{ColStart: 11, ColEnd: 19, Hex: "#f0f8ff"},
			&models.Colour{ColStart: 23, ColEnd: 26, Hex: "#0000ff"},
			&models.Colour{ColStart: 6, ColEnd: 9, Hex: "#808080"},
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
