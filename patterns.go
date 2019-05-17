package main

const (
	validHue     = `[0-9]{1,4}(?:\.[0-9]{1,2})?`
	num0to255    = `(?:2(?:[0-4][0-9]|5[0-5])|1?[0-9]?[0-9])(?:\.[0-9]{1,2})?`
	percentage   = `1?[0-9]{1,2}(?:\.[0-9]{1,2})?%`
	alphaPat     = `(?:0|1)?(?:\.[0-9]{1,2})?`
	hexDigit     = `[0-9a-fA-F]`
	rgbFuncParam = `(?:(?:` + num0to255 + `)|(?:` + percentage + `))`
)
