package parser

import (
	"github.com/rrethy/hexokinase/internal/models"
	"testing"
)

func TestParseRGBA(t *testing.T) {
	var tests = []struct {
		line string
		lnum int
		want colours
	}{
		{"rgba(0,0,0,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 13, Lnum: 0, Hex: "#000000"},
		}},
		{"rgba(176,253,35,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 18, Lnum: 0, Hex: "#b0fd23"},
		}},
		{"rgba(176,253,35,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 18, Lnum: 0, Hex: "#b0fd23"},
		}},

		// test various alphas
		{"rgba(176,253,35,1.0)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 20, Lnum: 0, Hex: "#b0fd23"},
		}},
		{"rgba(176,253,35,0)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 18, Lnum: 0, Hex: "#ffffff"},
		}},
		{"rgba(176,253,35,0.0)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 20, Lnum: 0, Hex: "#ffffff"},
		}},
		{"rgba(176,253,35,0.1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 20, Lnum: 0, Hex: "#f7fee9"},
		}},
		{"rgba(176,253,35,0.9)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 20, Lnum: 0, Hex: "#b7fd38"},
		}},

		// test percentages
		{"rgba(0%,253,35,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 17, Lnum: 0, Hex: "#00fd23"},
		}},
		{"rgba(100%,253,35,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 19, Lnum: 0, Hex: "#fffd23"},
		}},
		{"rgba(25%,253,35,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 18, Lnum: 0, Hex: "#3ffd23"},
		}},
		{"rgba(253,25%,35,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 18, Lnum: 0, Hex: "#fd3f23"},
		}},
		{"rgba(35,253,25%,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 18, Lnum: 0, Hex: "#23fd3f"},
		}},
		{"rgba(0%,25%,35%,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 18, Lnum: 0, Hex: "#003f59"},
		}},

		// test red value
		{"rgba(35,0,0,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 14, Lnum: 0, Hex: "#230000"},
		}},
		{"rgba(176,0,0,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 15, Lnum: 0, Hex: "#b00000"},
		}},
		{"rgba(215,0,0,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 15, Lnum: 0, Hex: "#d70000"},
		}},
		{"rgba(253,0,0,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 15, Lnum: 0, Hex: "#fd0000"},
		}},
		{"rgba(255,0,0,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 15, Lnum: 0, Hex: "#ff0000"},
		}},

		// test green value
		{"rgba(0,35,0,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 14, Lnum: 0, Hex: "#002300"},
		}},
		{"rgba(0,176,0,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 15, Lnum: 0, Hex: "#00b000"},
		}},
		{"rgba(0,215,0,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 15, Lnum: 0, Hex: "#00d700"},
		}},
		{"rgba(0,253,0,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 15, Lnum: 0, Hex: "#00fd00"},
		}},
		{"rgba(0,255,0,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 15, Lnum: 0, Hex: "#00ff00"},
		}},

		// test blue value
		{"rgba(0,0,35,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 14, Lnum: 0, Hex: "#000023"},
		}},
		{"rgba(0,0,176,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 15, Lnum: 0, Hex: "#0000b0"},
		}},
		{"rgba(0,0,215,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 15, Lnum: 0, Hex: "#0000d7"},
		}},
		{"rgba(0,0,253,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 15, Lnum: 0, Hex: "#0000fd"},
		}},
		{"rgba(0,0,255,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 15, Lnum: 0, Hex: "#0000ff"},
		}},

		// test multiple values
		{"rgba(0,0,255,1)rgba(176,253,35,1)  rgba(176,253,35,1)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 15, Lnum: 0, Hex: "#0000ff"},
			&models.Colour{ColStart: 16, ColEnd: 33, Lnum: 0, Hex: "#b0fd23"},
			&models.Colour{ColStart: 36, ColEnd: 53, Lnum: 0, Hex: "#b0fd23"},
		}},

		// tests invalid values
		{"rgba(256,0,0,1)", 0, []*models.Colour{}},
		{"rgba(0,0,256,1)", 0, []*models.Colour{}},
		{"rgba(0,0,256,1)", 0, []*models.Colour{}},
		{"rgba(1000,1000,1000,1)", 0, []*models.Colour{}},

		// test handling of whitespace
		{" rgba(0,0,0,1) ", 0, []*models.Colour{
			&models.Colour{ColStart: 2, ColEnd: 14, Lnum: 0, Hex: "#000000"},
		}},
		{" rgba(0,0,0,1) rgba(0,0,0,1) ", 0, []*models.Colour{
			&models.Colour{ColStart: 2, ColEnd: 14, Lnum: 0, Hex: "#000000"},
			&models.Colour{ColStart: 16, ColEnd: 28, Lnum: 0, Hex: "#000000"},
		}},
		{"rgba( 0 , 0 , 0 , 1 )", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 21, Lnum: 0, Hex: "#000000"},
		}},
		{"rgba(  0  ,  0  ,  0  ,  1  )", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 29, Lnum: 0, Hex: "#000000"},
		}},
		{"rgba(	0	,	0	,	0	,	1	)", 0, []*models.Colour{
			&models.Colour{ColStart: 1, ColEnd: 21, Lnum: 0, Hex: "#000000"},
		}},
	}
	for _, test := range tests {
		if got := parseRGBA(test.line, test.lnum); !areSameColours(got, test.want) {
			t.Errorf(" - %s\n\tgot:    %v\n\twanted: %v", test.line, got, test.want)
		}
	}
}
