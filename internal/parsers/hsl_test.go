package parser

import (
	"testing"
)

func TestParseHSL(t *testing.T) {
	var tests = []testData{}
	runTests("TestParseHSL", t, tests, parseHSL)
}
