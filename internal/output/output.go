package output

import (
	"fmt"
	"github.com/rrethy/hexokinase/internal/models"
	"os"
)

var (
	// Out TODO
	Out *os.File
)

// PrintColour TODO
func PrintColour(c *models.Colour) {
	fmt.Fprintf(Out, "%d:%d-%d:%s\n", c.Lnum, c.ColStart, c.ColEnd, c.Hex)
}
