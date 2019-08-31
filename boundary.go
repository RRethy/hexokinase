package main

var checkBoundary = false

// poor mans vimscript "\<\>"
func isWord(line string, start int, end int) bool {
	return (start == 0 || !isKeyword(line[start-1])) && (end == len(line) || !isKeyword(line[end]))
}

// TODO do a better job with utf-8
func isKeyword(c byte) bool {
	return (c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z') || (c >= '0' && c <= '9') || c == '_' || c == '-'
}
