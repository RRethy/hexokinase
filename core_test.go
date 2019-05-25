package main

import (
	"os"
	"testing"
)

func BenchmarkParseFile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		file, err := os.Open("./benchmark_colours.txt")
		if err != nil {
			b.Errorf("%v\n", err)
			continue
		}
		parseFile(file, "", -1)
		file.Close()
	}
}
