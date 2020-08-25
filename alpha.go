package main

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	bg = rgb{255, 255, 255}

	// ErrInvalidBgHex indicates that a bg value does not contain a valid six
	// digit hex
	ErrInvalidBgHex = errors.New("invalid hex value")

	// ErrInvalidBgRGB indicates that a bg value does not contain valid values
	// for r, g, b
	ErrInvalidBgRGB = errors.New("invalid r,g,b value")
)

// A BgError records an unsuccessful attempt to change the bg colour used to
// adjust colours with an alpha.
type BgError struct {
	Func  string // the failing function (SetBgHex, SetBgRGB, etc.)
	Value string // the bg value (either a hex or r:g:b)
	Err   error  // the reason for failure (e.g. ErrInvalidBgHex, ErrInvalidBgRGB, etc.)
}

func (e *BgError) Error() string {
	return "hexokinase." + e.Func + ": " + "setting bg " + e.Value + ": " + e.Err.Error()
}

// SetBgHex sets the bg hex to be used as the background for parsed colours.
// This affects rgba and hsla functions which are mixed with the bg based on
// their alpha.
func SetBgHex(hex string) error {
	const fnSetBgHex = "SetBgHex"

	bgHex := hexColour.FindString(hex)
	if len(bgHex) > 0 {
		bg = hexToRGB(bgHex)
		return nil
	}

	return &BgError{fnSetBgHex, hex, ErrInvalidBgHex}
}

// SetBgRGB sets the bg r, g, b to be used as the background for parsed
// colours.  This affects rgba and hsla functions which are mixed with the bg
// based on their alpha.
func SetBgRGB(r, g, b int) error {
	const fnSetBgRGB = "SetBgRGB"

	if r < 0 || r > 255 || g < 0 || g > 255 || b < 0 || b > 255 {
		return ErrInvalidBgRGB
	}

	bg = rgb{r, g, b}
	return nil
}

// withAlpha mixes r, g, b with bg based on alpha and returns the new r, g, b
// values
func withAlpha(r, g, b int, alpha float64) (int, int, int) {
	if alpha > 1.0 {
		alpha = 0
	}
	if alpha < 0.0 {
		alpha = 0
	}
	newR := int(float64(r)*alpha + float64(bg.r)*(1-alpha))
	newG := int(float64(g)*alpha + float64(bg.g)*(1-alpha))
	newB := int(float64(b)*alpha + float64(bg.b)*(1-alpha))
	return newR, newG, newB
}

func hexWithAlpha(hex string) string {
	hexAlpha, err := strconv.ParseInt(hex[6:8], 16, 32)
	alpha := 1.0
	if err != nil {
		alpha = 1
	} else {
		alpha = float64(hexAlpha) / 255
	}
	c := hexToRGB(fmt.Sprintf("#%s", hex[0:6]))
	return rgbToHex(withAlpha(c.r, c.g, c.b, alpha))
}
