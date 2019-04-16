package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) != 0 {
		for _, fname := range files {
			f, err := os.Open(fname)
			if err != nil {
				continue
			}
			parse(f)
			f.Close()
		}
	} else {
		parse(os.Stdin)
	}
}

func parse(f *os.File) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		fmt.Println(text)
	}
}
