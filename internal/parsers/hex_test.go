package parser

import (
	"github.com/rrethy/hexokinase/internal/models"
	"testing"
)

type colours []*models.Colour

func TestParseHex(t *testing.T) {
	var tests = []struct {
		line string
		want colours
	}{
		// test various values
		{" #fff ", []*models.Colour{
			&models.Colour{ColStart: 2, ColEnd: 5, Hex: "#ffffff"},
		}},
		{" #fff ", []*models.Colour{
			&models.Colour{ColStart: 2, ColEnd: 5, Hex: "#ffffff"},
		}},
		{" #FFF ", []*models.Colour{
			&models.Colour{ColStart: 2, ColEnd: 5, Hex: "#ffffff"},
		}},
		{" #ffffff ", []*models.Colour{
			&models.Colour{ColStart: 2, ColEnd: 8, Hex: "#ffffff"},
		}},
		{" #FFFFFF ", []*models.Colour{
			&models.Colour{ColStart: 2, ColEnd: 8, Hex: "#ffffff"},
		}},
		{"	#FFFFFF	", []*models.Colour{
			&models.Colour{ColStart: 2, ColEnd: 8, Hex: "#ffffff"},
		}},
		{" #A23f23 ", []*models.Colour{
			&models.Colour{ColStart: 2, ColEnd: 8, Hex: "#a23f23"},
		}},
		{"#a8f9e9", []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 7, Hex: "#a8f9e9"},
		}},

		// test invalid values
		{" # fff  ", []*models.Colour{}},
		{" #gggggg ", []*models.Colour{}},
		{" #banana ", []*models.Colour{}},
		{" banana ", []*models.Colour{}},
		{" #ggg ", []*models.Colour{}},

		// test multiple values
		{" #ae9 #A23f23 ", []*models.Colour{
			&models.Colour{ColStart: 2, ColEnd: 5, Lnum: 0, Hex: "#aaee99"},
			&models.Colour{ColStart: 7, ColEnd: 13, Lnum: 0, Hex: "#a23f23"},
		}},
		{"#ae9#A23f23", []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 4, Lnum: 0, Hex: "#aaee99"},
			&models.Colour{ColStart: 5, ColEnd: 11, Lnum: 0, Hex: "#a23f23"},
		}},
		{"#ae9#A23f23#000 #ae9#A23f23#000", []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 4, Lnum: 0, Hex: "#aaee99"},
			&models.Colour{ColStart: 5, ColEnd: 11, Lnum: 0, Hex: "#a23f23"},
			&models.Colour{ColStart: 12, ColEnd: 15, Lnum: 0, Hex: "#000000"},
			&models.Colour{ColStart: 17, ColEnd: 20, Lnum: 0, Hex: "#aaee99"},
			&models.Colour{ColStart: 21, ColEnd: 27, Lnum: 0, Hex: "#a23f23"},
			&models.Colour{ColStart: 28, ColEnd: 31, Lnum: 0, Hex: "#000000"},
		}},
	}
	for i, test := range tests {
		if got := parseHex(test.line); !areSameColours(got, test.want) {
			t.Errorf("test number: %d\n\tgot:    %v\n\twanted: %v", i, got, test.want)
		}
	}
}
