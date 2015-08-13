package main

import (
	"fmt"
	"os"
	"strings"
)

func wordToNumber(word string) string {
	var num string
	for _, rune := range strings.ToLower(word) {
		switch rune {
		case ' ':
			num += "0"
		case 'a', 'b', 'c':
			num += "2"
		case 'd', 'e', 'f':
			num += "3"
		case 'g', 'h', 'i':
			num += "4"
		case 'j', 'k', 'l':
			num += "5"
		case 'm', 'n', 'o':
			num += "6"
		case 'p', 'q', 'r', 's':
			num += "7"
		case 't', 'u', 'v':
			num += "8"
		case 'w', 'x', 'y', 'z':
			num += "9"
		}
	}

	return num
}

func main() {
	for _, arg := range os.Args[1:] {
		fmt.Println(wordToNumber(arg))
	}
}
