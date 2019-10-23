package main

import (
	"os" //library

	"github.com/01-edu/z01"
)

func main() {
	runes := []rune(os.Args[0])
	for _, letter := range runes {
		z01.PrintRune(letter)

	}
	z01.PrintRune('\n')
}
