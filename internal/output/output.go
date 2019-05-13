package output

import (
	"fmt"
	"github.com/rrethy/hexokinase/internal/models"
	"os"
)

var (
	// Out TODO
	out *os.File
)

// SetOut TODO
func SetOut(f *os.File) {
	out = f
}

// PrintColour TODO
func PrintColour(c *models.Colour) {
	fmt.Fprintf(out, "%d:%d-%d:%s\n", c.Lnum, c.ColStart, c.ColEnd, c.Hex)
}
