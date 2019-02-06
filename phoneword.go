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
		default:
			num += "*"
		case ' ', '0':
			num += "0"
		case '1':
			num += "1"
		case 'a', 'b', 'c', '2':
			num += "2"
		case 'd', 'e', 'f', '3':
			num += "3"
		case 'g', 'h', 'i', '4':
			num += "4"
		case 'j', 'k', 'l', '5':
			num += "5"
		case 'm', 'n', 'o', '6':
			num += "6"
		case 'p', 'q', 'r', 's', '7':
			num += "7"
		case 't', 'u', 'v', '8':
			num += "8"
		case 'w', 'x', 'y', 'z', '9':
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
