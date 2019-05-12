package parser

import (
	"fmt"
	"github.com/rrethy/hexokinase/internal/models"
	"strconv"
)

const (
	num0to255  = "2(?:[0-4][0-9]|5[0-5])|1?[0-9]?[0-9]"
	percentage = "1?[0-9]{0,2}%"
	funcParam  = "(?:(?:" + num0to255 + ")|(?:" + percentage + "))"
)

var (
	bgRGB = []int{255, 255, 255}
)

// SetBG TODO
func SetBG(hex string) {
	r, g, b := hexToRGB(hex)
	bgRGB = []int{r, g, b}
}

func areSameColours(colours1 []*models.Colour, colours2 []*models.Colour) bool {
	if len(colours1) != len(colours2) {
		return false
	}

	for i, colour1 := range colours1 {
		colour2 := colours2[i]
		if colour1.ColStart != colour2.ColStart ||
			colour1.ColEnd != colour2.ColEnd ||
			colour1.Lnum != colour2.Lnum ||
			colour1.Hex != colour2.Hex {
			return false
		}
	}

	return true
}

func strToDec(str string) (int, error) {
	if str[len(str)-1] == '%' {
		num, err := strconv.Atoi(str[:len(str)-1])
		if err != nil {
			return 0, err
		}
		return num * 255 / 100, nil
	}
	return strconv.Atoi(str)
}

func rgbToHex(r, g, b int) string {
	return fmt.Sprintf("#%02s%02s%02s",
		strconv.FormatInt(int64(r), 16),
		strconv.FormatInt(int64(g), 16),
		strconv.FormatInt(int64(b), 16))
}

func hexToRGB(hex string) (int, int, int) {
	r, err := strconv.ParseInt(hex[1:3], 16, 32)
	g, err := strconv.ParseInt(hex[1:3], 16, 32)
	b, err := strconv.ParseInt(hex[1:3], 16, 32)
	if err != nil {
		return bgRGB[0], bgRGB[1], bgRGB[2]
	}
	return int(r), int(g), int(b)
}

func setAlpha(r, g, b int, alpha float64) (int, int, int) {
	if alpha > 1.0 {
		alpha = 0
	}
	if alpha < 0.0 {
		alpha = 0
	}
	newR := int(float64(r)*alpha + float64(bgRGB[0])*(1.0-alpha))
	newG := int(float64(g)*alpha + float64(bgRGB[1])*(1.0-alpha))
	newB := int(float64(b)*alpha + float64(bgRGB[2])*(1.0-alpha))
	return newR, newG, newB
}
