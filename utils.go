package main

import (
	"fmt"
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

func percentageStrToInt(perStr string) (int, error) {
	return strconv.Atoi(perStr[:len(perStr)-1])
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

// https://github.com/lucasb-eyer/go-colorful/blob/30298f24079860c4dee452fdef6519b362a4a026/colors.go#L229
func hslToHex(h, s, l float64) string {
	if s == 0 {
		lInt := int(l * 255)
		return rgbToHex(lInt, lInt, lInt)
	}

	var r, g, b float64
	var t1 float64
	var t2 float64
	var tr float64
	var tg float64
	var tb float64

	if l < 0.5 {
		t1 = l * (1.0 + s)
	} else {
		t1 = l + s - l*s
	}

	t2 = 2*l - t1
	h = h / 360
	tr = h + 1.0/3.0
	tg = h
	tb = h - 1.0/3.0

	if tr < 0 {
		tr++
	}
	if tr > 1 {
		tr--
	}
	if tg < 0 {
		tg++
	}
	if tg > 1 {
		tg--
	}
	if tb < 0 {
		tb++
	}
	if tb > 1 {
		tb--
	}

	// Red
	if 6*tr < 1 {
		r = t2 + (t1-t2)*6*tr
	} else if 2*tr < 1 {
		r = t1
	} else if 3*tr < 2 {
		r = t2 + (t1-t2)*(2.0/3.0-tr)*6
	} else {
		r = t2
	}

	// Green
	if 6*tg < 1 {
		g = t2 + (t1-t2)*6*tg
	} else if 2*tg < 1 {
		g = t1
	} else if 3*tg < 2 {
		g = t2 + (t1-t2)*(2.0/3.0-tg)*6
	} else {
		g = t2
	}

	// Blue
	if 6*tb < 1 {
		b = t2 + (t1-t2)*6*tb
	} else if 2*tb < 1 {
		b = t1
	} else if 3*tb < 2 {
		b = t2 + (t1-t2)*(2.0/3.0-tb)*6
	} else {
		b = t2
	}

	rInt := int(r * 255)
	gInt := int(g * 255)
	bInt := int(b * 255)
	return rgbToHex(rInt, gInt, bInt)
}
