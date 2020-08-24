package main

import (
	"fmt"
	"strconv"
)

// toFullHex returns str if it is a six digit hex (#ffffff | 0xffffff),
// otherwise it will treat str as a three digit hex and convert it to a six
// digit hex.
// str MUST be either [#|0x]\x{3} or [#|0x]\x{6}.
func toFullHex(str string) string {
	switch len(str) {
	case 10:
		return hexWithAlpha(str[2:10])
	case 9:
		return hexWithAlpha(str[1:9])
	case 8:
		return fmt.Sprintf("#%c%c%c%c%c%c",
			str[2], str[3], str[4], str[5], str[6], str[7])
	case 7:
		return str
	case 5:
		return fmt.Sprintf("#%c%c%c%c%c%c",
			str[2], str[2], str[3], str[3], str[4], str[4])
	case 4:
		return fmt.Sprintf("#%c%c%c%c%c%c",
			str[1], str[1], str[2], str[2], str[3], str[3])
	default:
		return str
	}
}

// percentageStrToInt converts a str formatted as a percentage ("45.5%") into
// an int in [0, 100] where decimals are rounded down.
func percentageStrToInt(perStr string) (int, error) {
	num, err := strconv.ParseFloat(perStr[:len(perStr)-1], 64)
	return int(num), err
}

// strToDec converts a str formatted as either a percentage ("45.5%") or as an
// int ("35") into an int in [0, 255].
// "45.5%" would return 45 * 255 / 100 which is 114
// "35" would return 35
func strToDec(str string) (int, error) {
	if str[len(str)-1] == '%' {
		num, err := strconv.ParseFloat(str[:len(str)-1], 64)
		if err != nil {
			return 0, err
		}
		return int(num) * 255 / 100, nil
	}
	num, err := strconv.ParseFloat(str, 64)
	return int(num), err
}

// rgbaToHex returns a six digit hex value which represents r, g, b after it
// has been mixed with alpha.
// rgbaToHex(0, 0, 0, 0.0) returns "#000000"
// r, g, b must be in [0, 255]
// alpha must be in [0.0...1.0]
func rgbaToHex(r, g, b int, alpha float64) string {
	return rgbToHex(withAlpha(r, g, b, alpha))
}

// rgbToHex returns a six digit hex value which represents r, g, b
// rgbToHex(0, 0, 0) returns "#000000"
// r, g, b must be in [0, 255]
func rgbToHex(r, g, b int) string {
	return fmt.Sprintf("#%02s%02s%02s",
		strconv.FormatInt(int64(r), 16),
		strconv.FormatInt(int64(g), 16),
		strconv.FormatInt(int64(b), 16))
}

// hexToRGB returns the rgb representation of hex.
// hex MUST be #\x{6}
func hexToRGB(hex string) rgb {
	r, err := strconv.ParseInt(hex[1:3], 16, 32)
	g, err := strconv.ParseInt(hex[3:5], 16, 32)
	b, err := strconv.ParseInt(hex[5:7], 16, 32)
	if err != nil {
		return bg
	}
	return rgb{int(r), int(g), int(b)}
}

// hslaToHex returns a six digit hex value which represents h, s, l after it
// has been mixed with alpha.
// h must be [0, 359]
// s, l, alpha must be in [0.0...1.0]
func hslaToHex(h, s, l, alpha float64) string {
	r, g, b := hslToRGB(h, s, l)
	return rgbaToHex(r, g, b, alpha)
}

// hslToHex returns a six digit hex value which represents h, s, l
// h must be [0, 359]
// s, l must be in [0.0...1.0]
func hslToHex(h, s, l float64) string {
	r, g, b := hslToRGB(h, s, l)
	return rgbToHex(r, g, b)
}

// hslToRGB returns the rgb representation of h, s, l.
// h must be [0, 359]
// s, l must be in [0.0...1.0]
// https://github.com/lucasb-eyer/go-colorful/blob/30298f24079860c4dee452fdef6519b362a4a026/colors.go#L229
func hslToRGB(h, s, l float64) (int, int, int) {
	if s == 0 {
		lInt := int(l * 255)
		return lInt, lInt, lInt
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
	return rInt, gInt, bInt
}
