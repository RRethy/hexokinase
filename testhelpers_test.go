package main

import (
	"testing"
)

type testData struct {
	line string
	want colours
}

func runTests(tag string, t *testing.T, tests []testData, fun func(string) colours) {
	for _, test := range tests {
		if got := fun(test.line); !areSameColours(got, test.want) {
			t.Errorf(`
Func:   %+v
Input:  %s
Got:    %+v
Wanted: %+v`, tag, test.line, got, test.want)
		}
	}
}

func areSameColours(colours1 colours, colours2 colours) bool {
	if len(colours1) != len(colours2) {
		return false
	}

	for i, colour1 := range colours1 {
		colour2 := colours2[i]
		if colour1.ColStart != colour2.ColStart ||
			colour1.ColEnd != colour2.ColEnd ||
			colour1.Hex != colour2.Hex {
			return false
		}
	}

	return true
}
