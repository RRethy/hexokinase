package main

import (
	"os"
)

func main() {
	initializePalettes("./sample_palette.json")
	SetOut(os.Stdout)
	Read(os.Stdin)
}
