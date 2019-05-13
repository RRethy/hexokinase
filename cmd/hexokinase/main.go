package main

import (
	"github.com/rrethy/hexokinase/internal/output"
	"github.com/rrethy/hexokinase/internal/parsers"
	"os"
)

func main() {
	output.SetOut(os.Stdout)
	parser.Read(os.Stdin)
}
